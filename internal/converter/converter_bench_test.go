package converter

import (
    "testing"
    "image"
    "image/color"
    "github.com/hai119/Go-ASCII-generator/internal/testutils"
    "github.com/hai119/Go-ASCII-generator/internal/config"
    "fmt"
    "strings"
    "github.com/fogleman/gg"
    "runtime"
    "sync"
)

// 基准测试：亮度计算
func BenchmarkCalculateBrightness(b *testing.B) {
    sizes := []int{100, 500, 1000, 2000}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
            img := testutils.TestImage(size, size)
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                calculateBrightness(img, 0, 0, size/10, size/10)
            }
        })
    }
}

// 基准测试：平均颜色计算
func BenchmarkCalculateAverageColor(b *testing.B) {
    sizes := []int{100, 500, 1000, 2000}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
            img := testutils.TestImage(size, size)
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                calculateAverageColor(img, 0, 0, size/10, size/10)
            }
        })
    }
}

// 基准测试：图像转ASCII文本
func BenchmarkImageToText(b *testing.B) {
    sizes := []struct {
        width, height int
        name         string
    }{
        {640, 480, "small"},
        {1280, 720, "medium"},
        {1920, 1080, "large"},
    }

    for _, size := range sizes {
        b.Run(size.name, func(b *testing.B) {
            img := testutils.TestImage(size.width, size.height)
            cfg := &config.Config{
                NumCols:  100,
                CharMode: "complex",
            }
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                var buf strings.Builder
                processImageToText(img, cfg, &buf)
            }
        })
    }
}

// 基准测试：字符渲染
func BenchmarkCharacterRendering(b *testing.B) {
    sizes := []int{12, 16, 24, 32}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("font_size_%d", size), func(b *testing.B) {
            dc := gg.NewContext(size*2, size*2)
            if err := dc.LoadFontFace("../../fonts/DejaVuSansMono-Bold.ttf", float64(size)); err != nil {
                b.Fatal(err)
            }
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                dc.DrawString("@", float64(size/2), float64(size/2))
            }
        })
    }
}

// 基准测试：并行处理
func BenchmarkParallelProcessing(b *testing.B) {
    img := testutils.TestImage(1920, 1080)
    cfg := &config.Config{
        NumCols:  200,
        CharMode: "complex",
    }

    b.Run("sequential", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            processSequential(img, cfg)
        }
    })

    b.Run("parallel", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            processParallel(img, cfg)
        }
    })
}

// 辅助函数：顺序处理
func processSequential(img image.Image, cfg *config.Config) {
    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y
    cellWidth := float64(width) / float64(cfg.NumCols)
    cellHeight := 2 * cellWidth
    numRows := int(float64(height) / cellHeight)

    for i := 0; i < numRows; i++ {
        for j := 0; j < cfg.NumCols; j++ {
            calculateBrightness(img,
                int(float64(j)*cellWidth),
                int(float64(i)*cellHeight),
                int(cellWidth),
                int(cellHeight))
        }
    }
}

// 辅助函数：并行处理
func processParallel(img image.Image, cfg *config.Config) {
    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y
    cellWidth := float64(width) / float64(cfg.NumCols)
    cellHeight := 2 * cellWidth
    numRows := int(float64(height) / cellHeight)

    var wg sync.WaitGroup
    numWorkers := runtime.NumCPU()
    rowsPerWorker := numRows / numWorkers

    for w := 0; w < numWorkers; w++ {
        wg.Add(1)
        startRow := w * rowsPerWorker
        endRow := startRow + rowsPerWorker
        if w == numWorkers-1 {
            endRow = numRows
        }

        go func(start, end int) {
            defer wg.Done()
            for i := start; i < end; i++ {
                for j := 0; j < cfg.NumCols; j++ {
                    calculateBrightness(img,
                        int(float64(j)*cellWidth),
                        int(float64(i)*cellHeight),
                        int(cellWidth),
                        int(cellHeight))
                }
            }
        }(startRow, endRow)
    }

    wg.Wait()
}

// 基准测试：内存使用
func BenchmarkMemoryUsage(b *testing.B) {
    sizes := []struct {
        width, height int
        name         string
    }{
        {640, 480, "small"},
        {1280, 720, "medium"},
        {1920, 1080, "large"},
        {3840, 2160, "4k"},
    }

    for _, size := range sizes {
        b.Run(size.name, func(b *testing.B) {
            b.ReportAllocs()
            for i := 0; i < b.N; i++ {
                img := testutils.TestImage(size.width, size.height)
                cfg := &config.Config{
                    NumCols:  200,
                    CharMode: "complex",
                }
                var buf strings.Builder
                processImageToText(img, cfg, &buf)
            }
        })
    }
} 
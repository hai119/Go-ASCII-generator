package converter

import (
	"image"
	"image/color"
	"math"
	"testing"
)

// 模拟图像生成
func generateTestImage(width, height int, color color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color)
		}
	}
	return img
}

// 测试 WorkerPool
func testWorkerPool(t *testing.T) {
	// 测试用例：纯红色图像
	t.Run("TestPureRedImage", func(t *testing.T) {
		img := generateTestImage(100, 100, color.RGBA{255, 0, 0, 255}) // 纯红色图像
		pool := NewWorkerPool(4)
		pool.Start()

		cellWidth := 10
		cellHeight := 10
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				pool.jobs <- Job{
					img:        img,
					row:        row,
					col:        col,
					cellWidth:  cellWidth,
					cellHeight: cellHeight,
				}
			}
		}

		pool.Stop()

		// 验证结果
		for result := range pool.results {
			brightness := result.brightness
			if math.Abs(brightness-0.2989) > 0.01 {
				t.Errorf("Expected brightness around 0.2989 for red, got: %f", brightness)
			}
		}
	})

	// 测试用例：纯绿色图像
	t.Run("TestPureGreenImage", func(t *testing.T) {
		img := generateTestImage(100, 100, color.RGBA{0, 255, 0, 255}) // 纯绿色图像
		pool := NewWorkerPool(4)
		pool.Start()

		cellWidth := 10
		cellHeight := 10
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				pool.jobs <- Job{
					img:        img,
					row:        row,
					col:        col,
					cellWidth:  cellWidth,
					cellHeight: cellHeight,
				}
			}
		}

		pool.Stop()

		// 验证结果
		for result := range pool.results {
			brightness := result.brightness
			if math.Abs(brightness-0.5870) > 0.01 {
				t.Errorf("Expected brightness around 0.5870 for green, got: %f", brightness)
			}
		}
	})

	// 测试用例：纯蓝色图像
	t.Run("TestPureBlueImage", func(t *testing.T) {
		img := generateTestImage(100, 100, color.RGBA{0, 0, 255, 255}) // 纯蓝色图像
		pool := NewWorkerPool(4)
		pool.Start()

		cellWidth := 10
		cellHeight := 10
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				pool.jobs <- Job{
					img:        img,
					row:        row,
					col:        col,
					cellWidth:  cellWidth,
					cellHeight: cellHeight,
				}
			}
		}

		pool.Stop()

		// 验证结果
		for result := range pool.results {
			brightness := result.brightness
			if math.Abs(brightness-0.1140) > 0.01 {
				t.Errorf("Expected brightness around 0.1140 for blue, got: %f", brightness)
			}
		}
	})

	// 测试用例：渐变图像
	t.Run("TestGradientImage", func(t *testing.T) {
		img := image.NewRGBA(image.Rect(0, 0, 100, 100))
		for y := 0; y < 100; y++ {
			for x := 0; x < 100; x++ {
				r := uint8((float64(x) / 100) * 255)
				g := uint8((float64(y) / 100) * 255)
				b := uint8((float64(x+y) / 200) * 255)
				img.Set(x, y, color.RGBA{r, g, b, 255})
			}
		}

		pool := NewWorkerPool(4)
		pool.Start()

		cellWidth := 10
		cellHeight := 10
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				pool.jobs <- Job{
					img:        img,
					row:        row,
					col:        col,
					cellWidth:  cellWidth,
					cellHeight: cellHeight,
				}
			}
		}

		pool.Stop()

		// 验证结果
		for result := range pool.results {
			brightness := result.brightness
			if brightness < 0 || brightness > 1 {
				t.Errorf("Invalid brightness value: %f", brightness)
			}
		}
	})

	// 测试用例：无效图像路径
	t.Run("TestInvalidImage", func(t *testing.T) {
		img := generateTestImage(100, 100, color.RGBA{0, 0, 0, 255}) // 黑色图像
		pool := NewWorkerPool(4)
		pool.Start()

		// 使用不存在的图像路径
		img.(*image.RGBA).Set(101, 101, color.RGBA{255, 0, 0, 255}) // 超过图像范围，模拟错误

		pool.jobs <- Job{
			img:        img,
			row:        0,
			col:        0,
			cellWidth:  10,
			cellHeight: 10,
		}

		pool.Stop()

		// 验证结果
		for result := range pool.results {
			brightness := result.brightness
			if math.Abs(brightness) > 0.01 {
				t.Errorf("Expected error or invalid brightness, got: %f", brightness)
			}
		}
	})

	// 测试用例：空图像
	t.Run("TestEmptyImage", func(t *testing.T) {
		img := image.NewRGBA(image.Rect(0, 0, 0, 0)) // 空图像
		pool := NewWorkerPool(4)
		pool.Start()

		pool.jobs <- Job{
			img:        img,
			row:        0,
			col:        0,
			cellWidth:  10,
			cellHeight: 10,
		}

		pool.Stop()

		// 验证结果
		for result := range pool.results {
			if result.brightness != 0 {
				t.Errorf("Expected brightness to be 0 for empty image, got: %f", result.brightness)
			}
		}
	})

	// 测试用例：不同工作线程数量
	t.Run("TestMultipleWorkers", func(t *testing.T) {
		img := generateTestImage(100, 100, color.RGBA{255, 255, 255, 255}) // 白色图像

		for _, workers := range []int{1, 2, 4, 8} {
			pool := NewWorkerPool(workers)
			pool.Start()

			cellWidth := 10
			cellHeight := 10
			for row := 0; row < 10; row++ {
				for col := 0; col < 10; col++ {
					pool.jobs <- Job{
						img:        img,
						row:        row,
						col:        col,
						cellWidth:  cellWidth,
						cellHeight: cellHeight,
					}
				}
			}

			pool.Stop()

			// 验证结果
			for result := range pool.results {
				brightness := result.brightness
				if math.Abs(brightness-1) > 0.01 {
					t.Errorf("Expected brightness around 1 for white, got: %f", brightness)
				}
			}
		}
	})
}

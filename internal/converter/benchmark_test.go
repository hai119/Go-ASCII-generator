package converter

import (
    "image"
    "testing"

    "github.com/hai119/Go-ASCII-generator/internal/config"
)

func BenchmarkImageToText(b *testing.B) {
    cfg := &config.Config{
        InputPath:  "../../examples/input.jpg",
        OutputPath: "../../examples/output.txt",
        NumCols:    100,
        CharMode:   "simple",
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ImageToText(cfg)
    }
}

func BenchmarkCalculateBrightness(b *testing.B) {
    img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        calculateBrightness(img, 0, 0, 10, 10)
    }
}

func BenchmarkWorkerPool(b *testing.B) {
    pool := NewWorkerPool(4)
    img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
    
    b.ResetTimer()
    pool.Start()
    for i := 0; i < b.N; i++ {
        pool.jobs <- Job{
            img:        img,
            row:        i % 100,
            col:        i % 100,
            cellWidth:  10,
            cellHeight: 20,
        }
    }
    pool.Stop()
} 
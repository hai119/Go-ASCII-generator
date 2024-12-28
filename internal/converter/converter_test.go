package converter

import (
    "testing"
    "image"
    "image/color"
    "github.com/hai119/Go-ASCII-generator/internal/config"
    "github.com/hai119/Go-ASCII-generator/internal/testutils"
    "math"
    "strings"
)

func TestCalculateBrightness(t *testing.T) {
    tests := []struct {
        name     string
        img      image.Image
        x, y     int
        w, h     int
        expected float64
    }{
        {
            name:     "White image",
            img:      &image.Uniform{color.White},
            x:        0,
            y:        0,
            w:        10,
            h:        10,
            expected: 1.0,
        },
        {
            name:     "Black image",
            img:      &image.Uniform{color.Black},
            x:        0,
            y:        0,
            w:        10,
            h:        10,
            expected: 0.0,
        },
        // ... 更多测试用例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := calculateBrightness(tt.img, tt.x, tt.y, tt.w, tt.h)
            if got != tt.expected {
                t.Errorf("calculateBrightness() = %v, want %v", got, tt.expected)
            }
        })
    }
}

// 测试字符集选择
func TestGetCharList(t *testing.T) {
    tests := []struct {
        name     string
        mode     string
        expected int
    }{
        {
            name:     "Simple mode",
            mode:     "simple",
            expected: 10,
        },
        {
            name:     "Complex mode",
            mode:     "complex",
            expected: 70,
        },
        {
            name:     "Invalid mode",
            mode:     "invalid",
            expected: 70, // 默认使用复杂字符集
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            chars := getCharList(tt.mode)
            if len(chars) != tt.expected {
                t.Errorf("getCharList() returned %d chars, want %d", len(chars), tt.expected)
            }
        })
    }
}

// 测试颜色亮度计算
func TestCalculateColorBrightness(t *testing.T) {
    tests := []struct {
        name     string
        color    color.Color
        expected float64
    }{
        {
            name:     "White",
            color:    color.RGBA{255, 255, 255, 255},
            expected: 1.0,
        },
        {
            name:     "Black",
            color:    color.RGBA{0, 0, 0, 255},
            expected: 0.0,
        },
        {
            name:     "Red",
            color:    color.RGBA{255, 0, 0, 255},
            expected: 0.299,
        },
        {
            name:     "Green",
            color:    color.RGBA{0, 255, 0, 255},
            expected: 0.587,
        },
        {
            name:     "Blue",
            color:    color.RGBA{0, 0, 255, 255},
            expected: 0.114,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := calculateColorBrightness(tt.color)
            if math.Abs(got-tt.expected) > 0.001 {
                t.Errorf("calculateColorBrightness() = %v, want %v", got, tt.expected)
            }
        })
    }
}

// 测试平均颜色计算
func TestCalculateAverageColor(t *testing.T) {
    // 创建测试图像
    img := image.NewRGBA(image.Rect(0, 0, 2, 2))
    img.Set(0, 0, color.RGBA{255, 0, 0, 255})   // Red
    img.Set(1, 0, color.RGBA{0, 255, 0, 255})   // Green
    img.Set(0, 1, color.RGBA{0, 0, 255, 255})   // Blue
    img.Set(1, 1, color.RGBA{255, 255, 255, 255}) // White

    avgColor := calculateAverageColor(img, 0, 0, 2, 2)
    r, g, b, a := avgColor.RGBA()

    // 验证平均颜色
    expectedR := uint32(127.5 * 257) // 转换为16位色彩空间
    expectedG := uint32(127.5 * 257)
    expectedB := uint32(127.5 * 257)
    expectedA := uint32(255 * 257)

    if math.Abs(float64(r-expectedR)) > 257 || // 允许1的误差
        math.Abs(float64(g-expectedG)) > 257 ||
        math.Abs(float64(b-expectedB)) > 257 ||
        a != expectedA {
        t.Errorf("calculateAverageColor() = (%v, %v, %v, %v), want (%v, %v, %v, %v)",
            r, g, b, a, expectedR, expectedG, expectedB, expectedA)
    }
}

// 测试边界条件
func TestEdgeCases(t *testing.T) {
    tests := []struct {
        name string
        fn   func() error
    }{
        {
            name: "Empty image",
            fn: func() error {
                img := image.NewRGBA(image.Rect(0, 0, 0, 0))
                return processImageToText(img, &config.Config{NumCols: 100}, &strings.Builder{})
            },
        },
        {
            name: "Single pixel image",
            fn: func() error {
                img := image.NewRGBA(image.Rect(0, 0, 1, 1))
                return processImageToText(img, &config.Config{NumCols: 100}, &strings.Builder{})
            },
        },
        {
            name: "Very small number of columns",
            fn: func() error {
                img := testutils.TestImage(100, 100)
                return processImageToText(img, &config.Config{NumCols: 1}, &strings.Builder{})
            },
        },
        {
            name: "Very large number of columns",
            fn: func() error {
                img := testutils.TestImage(100, 100)
                return processImageToText(img, &config.Config{NumCols: 1000}, &strings.Builder{})
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := tt.fn(); err != nil {
                t.Errorf("Edge case %s failed: %v", tt.name, err)
            }
        })
    }
}

// ... 更多测试函数 
package converter

import (
    "image"
    "image/color"
    "testing"
)

func TestGetCharList(t *testing.T) {
    tests := []struct {
        name     string
        mode     string
        expected string
    }{
        {
            name:     "simple mode",
            mode:     "simple",
            expected: "@%#*+=-:. ",
        },
        {
            name:     "complex mode",
            mode:     "complex",
            expected: "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := getCharList(tt.mode)
            if result != tt.expected {
                t.Errorf("getCharList(%s) = %v, want %v", tt.mode, result, tt.expected)
            }
        })
    }
}

func TestCalculateBrightness(t *testing.T) {
    // 创建测试图像
    img := image.NewRGBA(image.Rect(0, 0, 2, 2))
    img.Set(0, 0, color.White)
    img.Set(1, 0, color.Black)
    img.Set(0, 1, color.Gray{Y: 128})
    img.Set(1, 1, color.White)

    tests := []struct {
        name     string
        x, y     int
        w, h     int
        expected float64
    }{
        {
            name:     "white pixel",
            x:        0,
            y:        0,
            w:        1,
            h:        1,
            expected: 1.0,
        },
        {
            name:     "black pixel",
            x:        1,
            y:        0,
            w:        1,
            h:        1,
            expected: 0.0,
        },
        {
            name:     "average of region",
            x:        0,
            y:        0,
            w:        2,
            h:        2,
            expected: 0.625, // (1.0 + 0.0 + 0.5 + 1.0) / 4
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := calculateBrightness(img, tt.x, tt.y, tt.w, tt.h)
            if result != tt.expected {
                t.Errorf("calculateBrightness() = %v, want %v", result, tt.expected)
            }
        })
    }
} 
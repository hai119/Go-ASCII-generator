package converter

import (
    "image"
    "image/color"
)

// MockImage 实现 image.Image 接口用于测试
type MockImage struct {
    width, height int
    pixels        [][]color.Color
}

func NewMockImage(width, height int) *MockImage {
    pixels := make([][]color.Color, height)
    for i := range pixels {
        pixels[i] = make([]color.Color, width)
        for j := range pixels[i] {
            pixels[i][j] = color.White
        }
    }
    return &MockImage{
        width:  width,
        height: height,
        pixels: pixels,
    }
}

func (m *MockImage) ColorModel() color.Model {
    return color.RGBAModel
}

func (m *MockImage) Bounds() image.Rectangle {
    return image.Rect(0, 0, m.width, m.height)
}

func (m *MockImage) At(x, y int) color.Color {
    if x < 0 || x >= m.width || y < 0 || y >= m.height {
        return color.Black
    }
    return m.pixels[y][x]
}

// SetPixel 设置指定位置的颜色
func (m *MockImage) SetPixel(x, y int, c color.Color) {
    if x >= 0 && x < m.width && y >= 0 && y < m.height {
        m.pixels[y][x] = c
    }
}

// ... 更多模拟对象和辅助��数 
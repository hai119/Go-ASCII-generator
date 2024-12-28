package testutils

import (
    "image"
    "image/color"
    "image/draw"
    "math/rand"
    "time"
)

// TestImage 生成测试用的图片
func TestImage(width, height int) image.Image {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    
    // 生成随机颜色
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            img.Set(x, y, color.RGBA{
                R: uint8(rand.Intn(256)),
                G: uint8(rand.Intn(256)),
                B: uint8(rand.Intn(256)),
                A: 255,
            })
        }
    }
    
    return img
}

// GenerateTestPattern 生成测试图案
func GenerateTestPattern(width, height int) image.Image {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    
    // 绘制棋盘格图案
    squareSize := 20
    for x := 0; x < width; x += squareSize {
        for y := 0; y < height; y += squareSize {
            col := color.RGBA{255, 255, 255, 255}
            if (x/squareSize+y/squareSize)%2 == 0 {
                col = color.RGBA{0, 0, 0, 255}
            }
            
            rect := image.Rect(x, y,
                min(x+squareSize, width),
                min(y+squareSize, height))
            draw.Draw(img, rect, &image.Uniform{col}, image.Point{}, draw.Src)
        }
    }
    
    return img
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// GenerateColorGradient 生成颜色渐变图案
func GenerateColorGradient(width, height int) image.Image {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            r := uint8(float64(x) / float64(width) * 255)
            g := uint8(float64(y) / float64(height) * 255)
            b := uint8((float64(x+y) / float64(width+height)) * 255)
            img.Set(x, y, color.RGBA{r, g, b, 255})
        }
    }
    
    return img
}

// GenerateNoise 生成噪点图案
func GenerateNoise(width, height int, density float64) image.Image {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    rand.Seed(time.Now().UnixNano())
    
    // 填充白色背景
    draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
    
    // 添加随机黑点
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            if rand.Float64() < density {
                img.Set(x, y, color.Black)
            }
        }
    }
    
    return img
}

// GenerateCircles 生成圆形图案
func GenerateCircles(width, height int, numCircles int) image.Image {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    rand.Seed(time.Now().UnixNano())
    
    // 填充白色背景
    draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
    
    // 绘制随机圆形
    for i := 0; i < numCircles; i++ {
        centerX := rand.Intn(width)
        centerY := rand.Intn(height)
        radius := rand.Intn(50) + 10
        col := color.RGBA{
            R: uint8(rand.Intn(256)),
            G: uint8(rand.Intn(256)),
            B: uint8(rand.Intn(256)),
            A: 255,
        }
        
        drawCircle(img, centerX, centerY, radius, col)
    }
    
    return img
}

// drawCircle 在图像上绘制圆形
func drawCircle(img *image.RGBA, centerX, centerY, radius int, col color.Color) {
    for x := centerX - radius; x <= centerX+radius; x++ {
        for y := centerY - radius; y <= centerY+radius; y++ {
            if x >= 0 && x < img.Bounds().Max.X &&
                y >= 0 && y < img.Bounds().Max.Y {
                dx := float64(x - centerX)
                dy := float64(y - centerY)
                if dx*dx+dy*dy <= float64(radius*radius) {
                    img.Set(x, y, col)
                }
            }
        }
    }
} 
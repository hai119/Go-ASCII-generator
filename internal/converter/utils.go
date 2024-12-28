package converter

import (
    "image"
    "image/color"
)

var (
    SimpleChars  = "@%#*+=-:. "
    ComplexChars = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
)

func getCharList(mode string) string {
    if mode == "simple" {
        return SimpleChars
    }
    return ComplexChars
}

func calculateBrightness(img image.Image, x, y, width, height int) float64 {
    var sum float64
    var count int
    
    for cy := y; cy < y+height && cy < img.Bounds().Max.Y; cy++ {
        for cx := x; cx < x+width && cx < img.Bounds().Max.X; cx++ {
            r, g, b, _ := img.At(cx, cy).RGBA()
            brightness := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535
            sum += brightness
            count++
        }
    }
    
    if count == 0 {
        return 0
    }
    return sum / float64(count)
}

func getBgColor(background string) color.Color {
    if background == "white" {
        return color.White
    }
    return color.Black
} 
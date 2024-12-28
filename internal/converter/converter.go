package converter

import (
	"image"
	"image/color"
)

var (
	SimpleChars  = "@%#*+=-:. "
	ComplexChars = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
)

// getCharList returns the character set based on mode
func getCharList(mode string) []rune {
	if mode == "simple" {
		return []rune(SimpleChars)
	}
	return []rune(ComplexChars)
}

// getBgColor returns the background color
func getBgColor(background string) color.Color {
	if background == "white" {
		return color.White
	}
	return color.Black
}

// calculateBrightness calculates the average brightness of an image region
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

func calculateAverageColor(img image.Image, x, y, width, height int) color.Color {
	var sumR, sumG, sumB uint32
	var count uint32

	for cy := y; cy < y+height && cy < img.Bounds().Max.Y; cy++ {
		for cx := x; cx < x+width && cx < img.Bounds().Max.X; cx++ {
			r, g, b, _ := img.At(cx, cy).RGBA()
			sumR += r
			sumG += g
			sumB += b
			count++
		}
	}

	if count == 0 {
		return color.RGBA{0, 0, 0, 255}
	}

	return color.RGBA{
		uint8(sumR / count >> 8),
		uint8(sumG / count >> 8),
		uint8(sumB / count >> 8),
		255,
	}
} 
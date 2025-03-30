package converter

import (
	"image"
	"image/color"
	"testing"
)

// MockImage is a helper function to generate a simple image for testing
func MockImageGeneration() image.Image {
	width, height := 3, 3
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Set some colors
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})     // Red
	img.Set(1, 0, color.RGBA{0, 255, 0, 255})     // Green
	img.Set(2, 0, color.RGBA{0, 0, 255, 255})     // Blue
	img.Set(0, 1, color.RGBA{255, 255, 0, 255})   // Yellow
	img.Set(1, 1, color.RGBA{0, 255, 255, 255})   // Cyan
	img.Set(2, 1, color.RGBA{255, 0, 255, 255})   // Magenta
	img.Set(0, 2, color.RGBA{0, 0, 0, 255})       // Black
	img.Set(1, 2, color.RGBA{255, 255, 255, 255}) // White
	img.Set(2, 2, color.RGBA{128, 128, 128, 255}) // Gray
	return img
}

func TestGetCharList(t *testing.T) {
	tests := []struct {
		mode     string
		expected []rune
	}{
		{"simple", []rune(SimpleChars)},
		{"complex", []rune(ComplexChars)},
	}

	for _, test := range tests {
		t.Run(test.mode, func(t *testing.T) {
			result := getCharList(test.mode)
			if len(result) != len(test.expected) {
				t.Errorf("expected length %d, got %d", len(test.expected), len(result))
			}
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("expected %c, got %c at index %d", test.expected[i], result[i], i)
				}
			}
		})
	}
}

func TestGetBgColor(t *testing.T) {
	tests := []struct {
		background string
		expected   color.Color
	}{
		{"white", color.White},
		{"black", color.Black},
	}

	for _, test := range tests {
		t.Run(test.background, func(t *testing.T) {
			result := getBgColor(test.background)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestCalculateAverageColor(t *testing.T) {
	img := MockImageGeneration()

	tests := []struct {
		x, y, width, height int
		expected            color.Color
	}{
		{0, 0, 1, 1, color.RGBA{255, 0, 0, 255}},     // Red pixel
		{1, 0, 1, 1, color.RGBA{0, 255, 0, 255}},     // Green pixel
		{0, 1, 1, 1, color.RGBA{255, 255, 0, 255}},   // Yellow pixel
		{0, 0, 3, 3, color.RGBA{128, 128, 128, 255}}, // Average of all colors in image
	}

	for _, test := range tests {
		t.Run("AverageColor", func(t *testing.T) {
			result := calculateAverageColor(img, test.x, test.y, test.width, test.height)
			r, g, b, a := result.RGBA()
			expectedR, expectedG, expectedB, expectedA := test.expected.RGBA()
			if r != expectedR || g != expectedG || b != expectedB || a != expectedA {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

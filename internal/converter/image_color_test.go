package converter

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"testing"

	"github.com/hai119/Go-ASCII-generator/internal/config"
	"github.com/stretchr/testify/assert"
)

// MockConfig is a helper function to create a mock configuration for testing.
func MockConfig(inputPath, outputPath string, numCols int, scale float64, charMode, background string) *config.Config {
	return &config.Config{
		InputPath:  inputPath,
		OutputPath: outputPath,
		NumCols:    numCols,
		Scale:      scale,
		CharMode:   charMode,
		Background: background,
	}
}

// MockImage is a helper function to generate a simple image for testing.
func MockImageGenerationV2() image.Image {
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

func TestImageToImageColor(t *testing.T) {
	// Create a mock config
	cfg := MockConfig("input.jpg", "output.jpg", 3, 1, "simple", "black")

	// Create a mock image
	img := MockImageGenerationV2()

	// Mock the os.Open function to simulate opening a file
	mockFile, err := os.Create("input.jpg")
	assert.NoError(t, err)
	defer mockFile.Close()

	err = jpeg.Encode(mockFile, img, nil)
	assert.NoError(t, err)

	// Test ImageToImageColor function
	t.Run("ImageToImageColor", func(t *testing.T) {
		// Mock the input path with an actual image
		err := ImageToImageColor(cfg)
		if err != nil {
			t.Errorf("expected no error, but got: %v", err)
		}
	})

	// Cleanup
	os.Remove("input.jpg")
}

// TestCalculateColorBrightness tests the calculateColorBrightness function
func TestCalculateColorBrightness(t *testing.T) {
	tests := []struct {
		color    color.Color
		expected float64
	}{
		{color.RGBA{255, 0, 0, 255}, 0.2126},   // Red
		{color.RGBA{0, 255, 0, 255}, 0.7152},   // Green
		{color.RGBA{0, 0, 255, 255}, 0.0722},   // Blue
		{color.RGBA{255, 255, 0, 255}, 0.9278}, // Yellow
		{color.RGBA{0, 0, 0, 255}, 0.0},        // Black
		{color.RGBA{255, 255, 255, 255}, 1.0},  // White
	}

	for _, test := range tests {
		t.Run("CalculateBrightness", func(t *testing.T) {
			result := calculateColorBrightness(test.color)
			assert.InDelta(t, test.expected, result, 0.2, "Brightness value is incorrect")
		})
	}
}

// TestHelper function to check if output file exists after ImageToImageColor run.
func TestOutputFileCreation(t *testing.T) {
	cfg := MockConfig("input.jpg", "output.jpg", 3, 1, "simple", "black")

	// Create a mock input file
	mockFile, err := os.Create("input.jpg")
	assert.NoError(t, err)
	defer mockFile.Close()

	// Create mock image and encode it to input.jpg
	img := MockImageGenerationV2()
	err = jpeg.Encode(mockFile, img, nil)
	assert.NoError(t, err)

	// Test ImageToImageColor function
	err = ImageToImageColor(cfg)
	assert.NoError(t, err)

	// Verify if output file exists
	_, err = os.Stat(cfg.OutputPath)
	assert.NoError(t, err, "Output file was not created")

	// Cleanup
	os.Remove("input.jpg")
	os.Remove(cfg.OutputPath)
}

// TestUnsupportedOutputFormat tests for unsupported output formats
func TestUnsupportedOutputFormat(t *testing.T) {
	cfg := MockConfig("input.jpg", "output.txt", 3, 1, "simple", "black")

	// Test ImageToImageColor function
	err := ImageToImageColor(cfg)
	assert.Error(t, err, "expected error for unsupported format")
}

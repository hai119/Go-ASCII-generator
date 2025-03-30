package converter

import (
	"image"
	"image/color"
	"math"
)

// Define sets of characters for different modes
var (
	SimpleChars  = "@%#*+=-:. "
	ComplexChars = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
)

// getCharList returns the character set based on the provided mode
// The mode can either be 'simple' or 'complex', and the corresponding character set will be returned.
func getCharList(mode string) []rune {
	// Add extra checks for mode variation
	if mode == "simple" {
		// Return the simple character set
		return []rune(SimpleChars)
	} else if mode == "complex" {
		// Return the complex character set
		return []rune(ComplexChars)
	}
	// Default case: return complex if mode is unrecognized
	return []rune(ComplexChars)
}

// getBgColor returns the background color based on the provided string
// It will return white for the string "white", and black for all other inputs.
func getBgColor(background string) color.Color {
	// Check if the background is 'white'
	if background == "white" {
		// Return white color
		return color.White
	}

	// Default case: return black for all other inputs
	return color.Black
}

// calculateBrightness calculates the average brightness of a region in the image
// This function iterates over the specified region defined by (x, y, width, height) and computes the average brightness
// based on the RGB values of the pixels in that region.
func calculateBrightness(img image.Image, x, y, width, height int) float64 {
	var sum float64
	var count int

	// Iterate through the specified region of the image
	for cy := y; cy < y+height && cy < img.Bounds().Max.Y; cy++ {
		for cx := x; cx < x+width && cx < img.Bounds().Max.X; cx++ {
			// Get the RGBA values of the current pixel
			r, g, b, _ := img.At(cx, cy).RGBA()
			// Calculate brightness using standard formula
			brightness := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535
			// Accumulate brightness values
			sum += brightness
			// Increment count of pixels
			count++
		}
	}

	// If no pixels were processed, return 0 brightness
	if count == 0 {
		return 0
	}

	// Calculate and return the average brightness of the region
	return sum / float64(count)
}

// calculateAverageColor computes the average color of a region in the image
// This function computes the average color by calculating the average RGBA values of all pixels in the region.
func calculateAverageColor(img image.Image, x, y, width, height int) color.Color {
	var sumR, sumG, sumB uint32
	var count uint32

	// Iterate through the image pixels within the specified region
	for cy := y; cy < y+height && cy < img.Bounds().Max.Y; cy++ {
		for cx := x; cx < x+width && cx < img.Bounds().Max.X; cx++ {
			// Get the RGBA values of the current pixel
			r, g, b, _ := img.At(cx, cy).RGBA()
			// Accumulate RGBA values
			sumR += r
			sumG += g
			sumB += b
			// Increment pixel count
			count++
		}
	}

	// If no pixels were processed, return black as the default color
	if count == 0 {
		return color.RGBA{0, 0, 0, 255}
	}

	// Calculate the average color by dividing the summed RGBA values by the count
	return color.RGBA{
		uint8(sumR / count >> 8),
		uint8(sumG / count >> 8),
		uint8(sumB / count >> 8),
		255, // Full opacity
	}
}

// calculateContrast computes the contrast of the image region based on its brightness
// This function calculates the contrast by computing the difference between the max and min brightness in the region
func calculateContrast(img image.Image, x, y, width, height int) float64 {
	var minBrightness, maxBrightness float64

	// Initialize minBrightness to a high value and maxBrightness to a low value
	minBrightness = math.MaxFloat64
	maxBrightness = -math.MaxFloat64

	// Iterate through the image region
	for cy := y; cy < y+height && cy < img.Bounds().Max.Y; cy++ {
		for cx := x; cx < x+width && cx < img.Bounds().Max.X; cx++ {
			// Get the brightness of the current pixel
			r, g, b, _ := img.At(cx, cy).RGBA()
			brightness := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535
			// Update the min and max brightness values
			if brightness < minBrightness {
				minBrightness = brightness
			}
			if brightness > maxBrightness {
				maxBrightness = brightness
			}
		}
	}

	// If no pixels were processed, return 0 contrast
	if minBrightness == maxBrightness {
		return 0
	}

	// Return the contrast as the difference between max and min brightness
	return maxBrightness - minBrightness
}

// calculateColorVariance computes the variance in color in a given region
// This function calculates the variance in color by comparing each pixel's RGB values
func calculateColorVariance(img image.Image, x, y, width, height int) float64 {
	var sumR, sumG, sumB float64
	var count float64

	// Calculate the sum of RGB values for each pixel in the region
	for cy := y; cy < y+height && cy < img.Bounds().Max.Y; cy++ {
		for cx := x; cx < x+width && cx < img.Bounds().Max.X; cx++ {
			r, g, b, _ := img.At(cx, cy).RGBA()
			sumR += float64(r)
			sumG += float64(g)
			sumB += float64(b)
			count++
		}
	}

	// Calculate the average RGB values
	avgR := sumR / count
	avgG := sumG / count
	avgB := sumB / count

	// Calculate the variance in the region
	var variance float64
	for cy := y; cy < y+height && cy < img.Bounds().Max.Y; cy++ {
		for cx := x; cx < x+width && cx < img.Bounds().Max.X; cx++ {
			r, g, b, _ := img.At(cx, cy).RGBA()
			variance += math.Pow(float64(r)-avgR, 2)
			variance += math.Pow(float64(g)-avgG, 2)
			variance += math.Pow(float64(b)-avgB, 2)
		}
	}

	// Return the calculated variance
	return variance / count
}

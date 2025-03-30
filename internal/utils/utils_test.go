package utils_test

import (
	"image"
	"image/color"
	"os"
	"testing"

	"github.com/hai119/Go-ASCII-generator/internal/utils"
)

func TestIsSupportedImageFormat(t *testing.T) {
	tests := []struct {
		filename string
		expected bool
	}{
		{"test.jpg", true},
		{"test.jpeg", true},
		{"test.png", true},
		{"test.gif", false},
		{"test.txt", false},
	}
	for _, tt := range tests {
		if got := utils.IsSupportedImageFormat(tt.filename); got != tt.expected {
			t.Errorf("IsSupportedImageFormat(%s) = %v; want %v", tt.filename, got, tt.expected)
		}
	}
}

func TestIsSupportedVideoFormat(t *testing.T) {
	tests := []struct {
		filename string
		expected bool
	}{
		{"video.mp4", true},
		{"video.avi", true},
		{"video.mov", true},
		{"video.mkv", false},
	}
	for _, tt := range tests {
		if got := utils.IsSupportedVideoFormat(tt.filename); got != tt.expected {
			t.Errorf("IsSupportedVideoFormat(%s) = %v; want %v", tt.filename, got, tt.expected)
		}
	}
}

func TestSaveImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})

	filename := "test_output.jpg"
	defer os.Remove(filename) // 清理测试文件

	if err := utils.SaveImage(img, filename); err != nil {
		t.Fatalf("SaveImage failed: %v", err)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("File %s was not created", filename)
	}
}

func TestEnsureDir(t *testing.T) {
	dir := "testdir"
	defer os.RemoveAll(dir) // 清理测试目录

	if err := utils.EnsureDir(dir); err != nil {
		t.Fatalf("EnsureDir failed: %v", err)
	}

	if info, err := os.Stat(dir); err != nil || !info.IsDir() {
		t.Errorf("Directory %s was not created", dir)
	}
}

func TestGetFileSize(t *testing.T) {
	filename := "testfile.txt"
	content := []byte("Hello, World!")
	if err := os.WriteFile(filename, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	size, err := utils.GetFileSize(filename)
	if err != nil {
		t.Fatalf("GetFileSize failed: %v", err)
	}

	if size != int64(len(content)) {
		t.Errorf("GetFileSize(%s) = %d; want %d", filename, size, len(content))
	}
}

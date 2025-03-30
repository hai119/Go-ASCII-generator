package converter

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strings"
	"testing"
)

// createTestImage 用于生成一个简单的测试图像并保存为JPEG格式
func createTestImage(path string) error {
	width, height := 100, 100
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 填充背景颜色为白色
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255}) // 设置为白色
		}
	}

	// 在中心画一个黑色的矩形
	for y := 40; y < 60; y++ {
		for x := 40; x < 60; x++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255}) // 设置为黑色
		}
	}

	// 保存图像为JPEG格式
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, nil)
}

// 测试 ImageToText 函数
func TestImageToText(t *testing.T) {
	// 创建一个临时图像文件作为输入
	inputPath := "test_input.jpg"
	outputPath := "test_output.txt"

	// 创建一个配置实例
	cfg := MockConfig(inputPath, outputPath, 10, 1, "simple", "black")

	// 模拟图像文件的生成
	err := createTestImage(inputPath)
	if err != nil {
		t.Fatalf("failed to create test image: %v", err)
	}
	defer os.Remove(inputPath)

	// 执行函数
	err = ImageToText(cfg)
	if err != nil {
		t.Fatalf("ImageToText failed: %v", err)
	}

	// 检查输出文件是否创建
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("output file does not exist")
	}

	// 清理
	os.Remove(inputPath)
	os.Remove(outputPath)
}

// 测试 ImageToImage 函数
func TestImageToImage(t *testing.T) {
	// 创建一个临时图像文件作为输入
	inputPath := "test_input.jpg"
	outputPath := "test_output.jpg"

	// 创建一个配置实例
	cfg := MockConfig(inputPath, outputPath, 10, 1, "simple", "white")

	// 模拟图像文件的生成
	err := createTestImage(inputPath)
	if err != nil {
		t.Fatalf("failed to create test image: %v", err)
	}
	defer os.Remove(inputPath)

	// 执行函数
	err = ImageToImage(cfg)
	if err != nil {
		t.Fatalf("ImageToImage failed: %v", err)
	}

	// 检查输出文件是否创建
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("output file does not exist")
	}

	// 检查是否是 JPEG 格式
	if !strings.HasSuffix(strings.ToLower(outputPath), ".jpg") && !strings.HasSuffix(strings.ToLower(outputPath), ".jpeg") {
		t.Fatalf("output file is not in JPEG format")
	}

	// 清理
	os.Remove(inputPath)
	os.Remove(outputPath)
}

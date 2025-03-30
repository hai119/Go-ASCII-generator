package converter

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
	"github.com/hai119/Go-ASCII-generator/internal/config"
)

// ImageToImageColor 转换图像为彩色ASCII艺术图像
func ImageToImageColor(cfg *config.Config) error {
	// 打开输入图像
	file, err := os.Open(cfg.InputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	// 解码图像
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	// 获取图像尺寸
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// 计算单元格大小
	cellWidth := float64(width) / float64(cfg.NumCols)
	cellHeight := 2 * cellWidth
	numRows := int(float64(height) / cellHeight)

	// 获取字符集
	chars := getCharList(cfg.CharMode)
	numChars := len(chars)

	// 创建输出图像
	dc := gg.NewContext(width, height)
	if cfg.Background == "white" {
		dc.SetRGB(1, 1, 1)
	} else {
		dc.SetRGB(0, 0, 0)
	}
	dc.Clear()

	// 设置字体
	if err := dc.LoadFontFace("fonts/DejaVuSansMono-Bold.ttf", 12*cfg.Scale); err != nil {
		return fmt.Errorf("failed to load font: %v", err)
	}

	// 转换图像为彩色ASCII艺术
	for i := 0; i < numRows; i++ {
		for j := 0; j < cfg.NumCols; j++ {
			// 计算当前单元格的平均颜色
			avgColor := calculateAverageColor(img,
				int(float64(j)*cellWidth),
				int(float64(i)*cellHeight),
				int(cellWidth),
				int(cellHeight))

			// 计算亮度并选择字符
			brightness := calculateColorBrightness(avgColor)
			charIndex := int(brightness * float64(numChars-1))
			if charIndex >= numChars {
				charIndex = numChars - 1
			}

			x := float64(j) * cellWidth
			y := float64(i)*cellHeight + cellHeight/2

			// 使用平均颜色绘制字符
			dc.SetColor(avgColor)
			dc.DrawStringAnchored(string(chars[charIndex]), x, y, 0, 0.5)
		}
	}

	// 保存输出图像
	outputDir := filepath.Dir(cfg.OutputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	output, err := os.Create(cfg.OutputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer output.Close()

	if strings.HasSuffix(strings.ToLower(cfg.OutputPath), ".jpg") ||
		strings.HasSuffix(strings.ToLower(cfg.OutputPath), ".jpeg") {
		return jpeg.Encode(output, dc.Image(), nil)
	}

	return fmt.Errorf("unsupported output format")
}

// 计算颜色亮度
func calculateColorBrightness(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	return (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535
}

func rotateImage(img image.Image, angle float64) image.Image {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	dc := gg.NewContext(width, height)
	dc.RotateAbout(angle, float64(width)/2, float64(height)/2)
	dc.DrawImage(img, 0, 0)
	return dc.Image()
}

func flipImage(img image.Image) image.Image {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	dc := gg.NewContext(width, height)
	dc.Scale(-1, 1)
	dc.DrawImage(img, -width, 0)
	return dc.Image()
}

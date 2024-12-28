package converter

import (
    "fmt"
    "image"
    "image/color"
    "image/jpeg"
    "os"
    "strings"

    "github.com/fogleman/gg"
    "github.com/hai119/Go-ASCII-generator/internal/config"
)

func ImageToText(cfg *config.Config) error {
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

    // 创建输出文件
    output, err := os.Create(cfg.OutputPath)
    if err != nil {
        return fmt.Errorf("failed to create output file: %v", err)
    }
    defer output.Close()

    // 转换图像为ASCII文本
    for i := 0; i < numRows; i++ {
        for j := 0; j < cfg.NumCols; j++ {
            brightness := calculateBrightness(img, 
                int(float64(j)*cellWidth), 
                int(float64(i)*cellHeight),
                int(cellWidth),
                int(cellHeight))
            
            charIndex := int(brightness * float64(numChars-1))
            if charIndex >= numChars {
                charIndex = numChars - 1
            }
            
            fmt.Fprint(output, string(chars[charIndex]))
        }
        fmt.Fprintln(output)
    }

    return nil
}

func ImageToImage(cfg *config.Config) error {
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
    dc.SetColor(getBgColor(cfg.Background))
    dc.Clear()

    // 设置字体
    if err := dc.LoadFontFace("fonts/DejaVuSansMono-Bold.ttf", 12*cfg.Scale); err != nil {
        return fmt.Errorf("failed to load font: %v", err)
    }

    // 设置前景色
    if cfg.Background == "white" {
        dc.SetColor(color.Black)
    } else {
        dc.SetColor(color.White)
    }

    // 转换图像为ASCII艺术
    for i := 0; i < numRows; i++ {
        for j := 0; j < cfg.NumCols; j++ {
            brightness := calculateBrightness(img,
                int(float64(j)*cellWidth),
                int(float64(i)*cellHeight),
                int(cellWidth),
                int(cellHeight))
            
            charIndex := int(brightness * float64(numChars-1))
            if charIndex >= numChars {
                charIndex = numChars - 1
            }
            
            x := float64(j) * cellWidth
            y := float64(i) * cellHeight + cellHeight/2
            
            dc.DrawStringAnchored(string(chars[charIndex]), x, y, 0, 0.5)
        }
    }

    // 保存输出图像
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
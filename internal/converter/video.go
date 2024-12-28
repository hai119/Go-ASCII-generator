package converter

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "image/jpeg"
    "io/ioutil"

    "github.com/hai119/Go-ASCII-generator/internal/config"
)

// VideoToText converts video to ASCII text
func VideoToText(cfg *config.Config) error {
    // 创建临时目录存放帧
    tempDir, err := ioutil.TempDir("", "ascii-frames-")
    if err != nil {
        return fmt.Errorf("failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    // 使用ffmpeg提取帧
    framePattern := filepath.Join(tempDir, "frame-%d.jpg")
    cmd := exec.Command("ffmpeg", "-i", cfg.InputPath, "-vf", "fps=10", framePattern)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract frames: %v", err)
    }

    // 创建输出文件
    outputDir := filepath.Dir(cfg.OutputPath)
    if err := os.MkdirAll(outputDir, 0755); err != nil {
        return fmt.Errorf("failed to create output directory: %v", err)
    }

    output, err := os.Create(cfg.OutputPath)
    if err != nil {
        return fmt.Errorf("failed to create output file: %v", err)
    }
    defer output.Close()

    // 处理每一帧
    frameFiles, err := filepath.Glob(filepath.Join(tempDir, "frame-*.jpg"))
    if err != nil {
        return fmt.Errorf("failed to list frames: %v", err)
    }

    for frameNum, framePath := range frameFiles {
        // 读取帧
        file, err := os.Open(framePath)
        if err != nil {
            return fmt.Errorf("failed to open frame: %v", err)
        }

        img, err := jpeg.Decode(file)
        file.Close()
        if err != nil {
            return fmt.Errorf("failed to decode frame: %v", err)
        }

        // 获取帧尺寸
        bounds := img.Bounds()
        width, height := bounds.Max.X, bounds.Max.Y

        // 计算单元格大小
        cellWidth := float64(width) / float64(cfg.NumCols)
        cellHeight := 2 * cellWidth
        numRows := int(float64(height) / cellHeight)

        // 获取字符集
        chars := getCharList(cfg.CharMode)
        numChars := len(chars)

        // 生成ASCII帧
        var frameText strings.Builder
        frameText.WriteString(fmt.Sprintf("Frame %d:\n", frameNum))

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

                frameText.WriteString(string(chars[charIndex]))
            }
            frameText.WriteString("\n")
        }
        frameText.WriteString("\n")

        // 写入输出文件
        if _, err := output.WriteString(frameText.String()); err != nil {
            return fmt.Errorf("failed to write frame: %v", err)
        }
    }

    return nil
} 
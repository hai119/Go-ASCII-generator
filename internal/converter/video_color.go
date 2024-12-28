package converter

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "image/jpeg"
    "io/ioutil"

    "github.com/fogleman/gg"
    "github.com/hai119/Go-ASCII-generator/internal/config"
)

func VideoToVideoColor(cfg *config.Config) error {
    // 创建临时目录
    tempDir, err := ioutil.TempDir("", "ascii-frames-")
    if err != nil {
        return fmt.Errorf("failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    // 提取原始帧
    inputFramePattern := filepath.Join(tempDir, "input-frame-%d.jpg")
    cmd := exec.Command("ffmpeg", "-i", cfg.InputPath, "-vf", "fps=10", inputFramePattern)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract frames: %v", err)
    }

    // 创建输出目录
    outputFrameDir := filepath.Join(tempDir, "output-frames")
    if err := os.MkdirAll(outputFrameDir, 0755); err != nil {
        return fmt.Errorf("failed to create output frame directory: %v", err)
    }

    // 处理每一帧
    frameFiles, err := filepath.Glob(filepath.Join(tempDir, "input-frame-*.jpg"))
    if err != nil {
        return fmt.Errorf("failed to list frames: %v", err)
    }

    for _, framePath := range frameFiles {
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

        // 处理帧
        bounds := img.Bounds()
        width, height := bounds.Max.X, bounds.Max.Y

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

        // 计算单元格大小
        cellWidth := float64(width) / float64(cfg.NumCols)
        cellHeight := 2 * cellWidth
        numRows := int(float64(height) / cellHeight)

        // 获取字符集
        chars := getCharList(cfg.CharMode)
        numChars := len(chars)

        // 转换为ASCII艺术
        for i := 0; i < numRows; i++ {
            for j := 0; j < cfg.NumCols; j++ {
                avgColor := calculateAverageColor(img,
                    int(float64(j)*cellWidth),
                    int(float64(i)*cellHeight),
                    int(cellWidth),
                    int(cellHeight))

                brightness := calculateColorBrightness(avgColor)
                charIndex := int(brightness * float64(numChars-1))
                if charIndex >= numChars {
                    charIndex = numChars - 1
                }

                x := float64(j) * cellWidth
                y := float64(i) * cellHeight + cellHeight/2

                dc.SetColor(avgColor)
                dc.DrawStringAnchored(string(chars[charIndex]), x, y, 0, 0.5)
            }
        }

        // 保存处理后的帧
        outputFramePath := filepath.Join(outputFrameDir, filepath.Base(framePath))
        outFile, err := os.Create(outputFramePath)
        if err != nil {
            return fmt.Errorf("failed to create output frame: %v", err)
        }

        if err := jpeg.Encode(outFile, dc.Image(), nil); err != nil {
            outFile.Close()
            return fmt.Errorf("failed to encode output frame: %v", err)
        }
        outFile.Close()
    }

    // 合成视频
    outputFramePattern := filepath.Join(outputFrameDir, "input-frame-%d.jpg")
    cmd = exec.Command("ffmpeg",
        "-framerate", "10",
        "-i", outputFramePattern,
        "-c:v", "libx264",
        "-pix_fmt", "yuv420p",
        cfg.OutputPath)
    
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to create output video: %v", err)
    }

    return nil
} 
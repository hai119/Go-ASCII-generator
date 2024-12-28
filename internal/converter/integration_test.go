package converter

import (
    "testing"
    "os"
    "path/filepath"
    "github.com/hai119/Go-ASCII-generator/internal/config"
    "github.com/hai119/Go-ASCII-generator/internal/testutils"
)

func TestImageToTextIntegration(t *testing.T) {
    // 创建临时目录
    tempDir, err := os.MkdirTemp("", "ascii-test-")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(tempDir)
    
    // 生成测试图像
    testCases := []struct {
        name     string
        genFunc  func(int, int) image.Image
        width    int
        height   int
        expected string
    }{
        {
            name:    "Checkerboard pattern",
            genFunc: testutils.GenerateTestPattern,
            width:   100,
            height:  100,
        },
        {
            name:    "Color gradient",
            genFunc: testutils.GenerateColorGradient,
            width:   100,
            height:  100,
        },
        {
            name: "Random noise",
            genFunc: func(w, h int) image.Image {
                return testutils.GenerateNoise(w, h, 0.3)
            },
            width:  100,
            height: 100,
        },
        // ... 更多测试用例
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // 生成输入图像
            inputPath := filepath.Join(tempDir, "input.jpg")
            outputPath := filepath.Join(tempDir, "output.txt")
            
            img := tc.genFunc(tc.width, tc.height)
            
            // 保存测试图像
            f, err := os.Create(inputPath)
            if err != nil {
                t.Fatal(err)
            }
            if err := jpeg.Encode(f, img, nil); err != nil {
                f.Close()
                t.Fatal(err)
            }
            f.Close()
            
            // 创建配置
            cfg := &config.Config{
                InputPath:  inputPath,
                OutputPath: outputPath,
                Mode:      "image2text",
                NumCols:   50,
                CharMode:  "complex",
            }
            
            // 执行转换
            if err := ImageToText(cfg); err != nil {
                t.Fatal(err)
            }
            
            // 验证输出文件存在
            if _, err := os.Stat(outputPath); os.IsNotExist(err) {
                t.Error("Output file was not created")
            }
            
            // 读取并验证输出内容
            content, err := os.ReadFile(outputPath)
            if err != nil {
                t.Fatal(err)
            }
            
            if len(content) == 0 {
                t.Error("Output file is empty")
            }
            
            // 可以添加更多具体的输出验证
        })
    }
}

// ... 更多集成测试 
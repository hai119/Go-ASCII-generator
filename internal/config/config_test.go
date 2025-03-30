package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseFlags(t *testing.T) {
	// 设置模拟的命令行参数
	os.Args = []string{"program", "-input", "data/input.jpg", "-output", "data/output.txt", "-mode", "image2text", "-cols", "120", "-bg", "white", "-char-mode", "simple", "-scale", "2.0", "-fps", "30", "-overlay", "0.5", "-lang", "chinese"}

	// 调用 ParseFlags 进行解析
	cfg := ParseFlags()

	// 检查解析后的配置是否正确
	absInputPath, _ := filepath.Abs("data/input.jpg")
	absOutputPath, _ := filepath.Abs("data/output.txt")
	assert.Equal(t, absInputPath, cfg.InputPath)
	assert.Equal(t, absOutputPath, cfg.OutputPath)
	assert.Equal(t, "image2text", cfg.Mode)
	assert.Equal(t, 120, cfg.NumCols)
	assert.Equal(t, "white", cfg.Background)
	assert.Equal(t, "simple", cfg.CharMode)
	assert.Equal(t, 2.0, cfg.Scale)
	assert.Equal(t, 30, cfg.FPS)
	assert.Equal(t, 0.5, cfg.OverlayRatio)
	assert.Equal(t, "chinese", cfg.Language)
}

func TestProcessPaths(t *testing.T) {
	// 设置模拟的配置
	cfg := &Config{
		InputPath:  "data/input.jpg",
		OutputPath: "data/output.txt",
	}

	// 调用 ProcessPaths 来处理路径
	cfg.ProcessPaths()

	// 验证输入路径和输出路径的处理结果
	absInputPath, _ := os.Getwd()
	absInputPath = filepath.Join(absInputPath, "data/input.jpg")
	assert.Equal(t, absInputPath, cfg.InputPath)

	absOutputPath, _ := os.Getwd()
	absOutputPath = filepath.Join(absOutputPath, "data/output.txt")
	assert.Equal(t, absOutputPath, cfg.OutputPath)

	// 检查输出目录是否存在
	outputDir := filepath.Dir(absOutputPath)
	_, err := os.Stat(outputDir)
	assert.False(t, os.IsNotExist(err), "Output directory should exist")
}

func TestRetryOperation(t *testing.T) {
	// 模拟一个失败的操作
	failOperation := func() error {
		return fmt.Errorf("operation failed")
	}

	// 测试重试操作，最多重试3次，间隔1秒
	err := RetryOperation(failOperation, 3, 1*time.Second)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "operation failed after 3 retries")

	// 模拟一个成功的操作
	successOperation := func() error {
		return nil
	}

	// 测试成功操作
	err = RetryOperation(successOperation, 3, 1*time.Second)
	assert.NoError(t, err)
}

func TestIsVerboseMode(t *testing.T) {
	// 设置环境变量为 "true"
	os.Setenv("VERBOSE_MODE", "true")
	assert.True(t, isVerboseMode(), "Verbose mode should be enabled")

	// 设置环境变量为 "false"
	os.Setenv("VERBOSE_MODE", "false")
	assert.False(t, isVerboseMode(), "Verbose mode should be disabled")

	// 设置环境变量为空
	os.Setenv("VERBOSE_MODE", "")
	assert.False(t, isVerboseMode(), "Verbose mode should be disabled")
}

func TestIsValidMode(t *testing.T) {
	// 验证有效模式
	assert.True(t, isValidMode("image2text"), "Mode 'image2text' should be valid")
	assert.True(t, isValidMode("image2image"), "Mode 'image2image' should be valid")
	assert.True(t, isValidMode("video2video"), "Mode 'video2video' should be valid")

	// 验证无效模式
	assert.False(t, isValidMode("invalidMode"), "Mode 'invalidMode' should be invalid")
}

func TestPrintConfig(t *testing.T) {
	// 创建一个配置实例
	cfg := &Config{
		InputPath:    "data/input.jpg",
		OutputPath:   "data/output.txt",
		Mode:         "image2text",
		NumCols:      100,
		Background:   "black",
		CharMode:     "complex",
		Scale:        1.0,
		FPS:          30,
		OverlayRatio: 0.5,
		Language:     "english",
	}

	// 由于是打印输出，我们这里不直接断言，而是通过捕获输出验证
	// 使用 `os.Stdout` 来捕获输出
	// 使用自定义的函数来验证输出
	printConfig(cfg)
	// 可以扩展捕获输出进行断言
}

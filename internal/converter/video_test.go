package converter

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hai119/Go-ASCII-generator/internal/config"
)

// 测试 VideoToText
func TestVideoToText(t *testing.T) {
	// 创建一个临时的伪视频文件
	inputPath := filepath.Join(t.TempDir(), "fake_video.mp4")
	if err := createFakeVideoFile(inputPath); err != nil {
		t.Fatalf("failed to create fake video: %v", err)
	}

	// 设置测试配置
	cfg := &config.Config{
		InputPath:  inputPath,
		OutputPath: filepath.Join(t.TempDir(), "output.txt"),
		NumCols:    10,
		CharMode:   "standard",
		Scale:      1,
		Background: "white",
	}

	// 执行 VideoToText 函数
	err := VideoToText(cfg)
	if err != nil {
		t.Fatalf("VideoToText failed: %v", err)
	}

	// 检查输出文件是否存在并包含预期内容
	output, err := ioutil.ReadFile(cfg.OutputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	// 检查文件是否包含帧内容
	if !strings.Contains(string(output), "Frame") {
		t.Fatalf("expected 'Frame' in the output, but got: %s", string(output))
	}

	// 检查帧文本是否按预期格式生成
	if len(output) == 0 {
		t.Fatal("output file is empty")
	}
}

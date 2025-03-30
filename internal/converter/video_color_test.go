package converter

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/hai119/Go-ASCII-generator/internal/config"
)

// 为了模拟 exec.Command 的行为，我们使用一个可替换的函数
var execCommand = exec.Command

// 为了模拟 ioutil.TempDir 的行为，我们使用一个可替换的函数
var ioutilTempDir = os.MkdirTemp

// Mock exec.Command to simulate success (skip actual ffmpeg execution)
func mockExecCommand(name string, arg ...string) *exec.Cmd {
	if name == "ffmpeg" {
		// 模拟 ffmpeg 成功的命令执行，不做任何实际操作
		return &exec.Cmd{
			ProcessState: &os.ProcessState{},
		}
	}
	// 返回默认的 cmd，如果需要支持其他命令
	return exec.Command(name, arg...)
}

// 创建一个假的视频文件，用于测试
func createFakeVideoFile(filePath string) error {
	// 使用正确的 ffmpeg 命令：指定分辨率 (s) 和帧率 (r)
	cmd := exec.Command("ffmpeg", "-f", "lavfi", "-t", "1", "-i", "testsrc=s=1920x1080:r=30", filePath)

	// 捕获 stderr 输出
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		// 输出 stderr 内容，以便调试
		return fmt.Errorf("failed to create fake video file: %v, stderr: %s", err, stderr.String())
	}

	return nil
}

// Test VideoToVideoColor function
func TestVideoToVideoColor(t *testing.T) {
	// 创建临时目录
	tempDir := t.TempDir()

	// 创建一个假的视频文件
	inputPath := filepath.Join(tempDir, "test_video.mp4")
	if err := createFakeVideoFile(inputPath); err != nil {
		t.Fatalf("failed to create fake video file: %v", err)
	}

	cfg := &config.Config{
		InputPath:  inputPath,                                  // 使用创建的假视频文件
		OutputPath: filepath.Join(tempDir, "output_video.mp4"), // 输出路径也是假的
		Background: "black",
		Scale:      1,
		NumCols:    50,
		CharMode:   "default",
	}

	// 将 execCommand 和 ioutilTempDir 替换为我们定义的 mock 函数
	execCommand = mockExecCommand

	// 测试成功场景
	err := VideoToVideoColor(cfg)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

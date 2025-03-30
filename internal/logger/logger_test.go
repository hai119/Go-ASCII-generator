package logger_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/hai119/Go-ASCII-generator/internal/logger"
)

func TestLogging(t *testing.T) {
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		t.Fatalf("Failed to create log directory: %v", err)
	}

	logFile := filepath.Join(logDir, "app_"+time.Now().Format("2006-01-02")+".log")
	defer os.RemoveAll(logDir) // 清理日志目录

	logger.Info("This is an info message")
	logger.Error("This is an error message")
	logger.Debug("This is a debug message")

	// 等待日志写入
	time.Sleep(100 * time.Millisecond)

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Fatalf("Log file %s was not created", logFile)
	}

	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	logContent := string(content)
	t.Logf("Log file content:\n%s", logContent) // 打印日志内容进行调试

	expectedKeywords := []string{"INFO:", "ERROR:", "DEBUG:"}
	for _, keyword := range expectedKeywords {
		if !strings.Contains(logContent, keyword) {
			t.Errorf("Log file does not contain expected keyword: %s", keyword)
		}
	}
}

func contains(content, substr string) bool {
	return len(content) >= len(substr) &&
		(content == substr || contains(content[1:], substr))
}

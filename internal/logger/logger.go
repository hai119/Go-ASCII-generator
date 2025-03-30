package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
)

func init() {
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile(
		filepath.Join(logDir, fmt.Sprintf("app_%s.log", time.Now().Format("2006-01-02"))),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	infoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
}

// Info 记录信息日志
func Info(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

// Error 记录错误日志
func Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	errorLogger.Printf("%s:%d: "+format, append([]interface{}{file, line}, v...)...)
}

// Debug 记录调试日志
func Debug(format string, v ...interface{}) {
	debugLogger.Printf(format, v...)
}

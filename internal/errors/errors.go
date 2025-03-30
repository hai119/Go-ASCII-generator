package errors

import "fmt"

// Error types
var (
	ErrInvalidInput    = fmt.Errorf("invalid input")
	ErrInvalidOutput   = fmt.Errorf("invalid output")
	ErrUnsupportedMode = fmt.Errorf("unsupported mode")
	ErrFontNotFound    = fmt.Errorf("font not found")
	ErrProcessing      = fmt.Errorf("processing error")
	ErrDatabase        = fmt.Errorf("database error")
	ErrFileNotFound    = fmt.Errorf("file not found")
	ErrNetworkIssue    = fmt.Errorf("network issue")
	ErrTimeout         = fmt.Errorf("timeout")
	ErrUnauthorized    = fmt.Errorf("unauthorized access")
	ErrInternal        = fmt.Errorf("internal server error")
	ErrInvalidRequest  = fmt.Errorf("invalid request")
	ErrNotImplemented  = fmt.Errorf("not implemented")
)

// AppError 应用错误结构
type AppError struct {
	Err         error  // 原始错误
	Message     string // 错误消息
	Code        int    // 错误代码
	Details     string // 错误详情，用于提供更详细的错误信息
	Timestamp   string // 错误发生的时间戳
	UserContext string // 当前用户上下文
}

// NewAppError 创建新的应用错误
func NewAppError(err error, message string, code int, details, timestamp, userContext string) *AppError {
	return &AppError{
		Err:         err,
		Message:     message,
		Code:        code,
		Details:     details,
		Timestamp:   timestamp,
		UserContext: userContext,
	}
}

// Error 实现 error 接口
func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s: %v (Code: %d) - Details: %s - User: %s", e.Timestamp, e.Message, e.Err, e.Code, e.Details, e.UserContext)
}

// WrapError 包装错误，保持错误链
func WrapError(err error, message string, details, timestamp, userContext string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("[%s] %s: %w - Details: %s - User: %s", timestamp, message, err, details, userContext)
}

// NewDatabaseError 创建数据库相关的应用错误
func NewDatabaseError(err error, message string, code int, details, timestamp, userContext string) *AppError {
	return &AppError{
		Err:         err,
		Message:     message,
		Code:        code,
		Details:     details,
		Timestamp:   timestamp,
		UserContext: userContext,
	}
}

// NewFileError 创建文件相关的应用错误
func NewFileError(err error, message string, code int, details, timestamp, userContext string) *AppError {
	return &AppError{
		Err:         err,
		Message:     message,
		Code:        code,
		Details:     details,
		Timestamp:   timestamp,
		UserContext: userContext,
	}
}

// NewNetworkError 创建网络相关的应用错误
func NewNetworkError(err error, message string, code int, details, timestamp, userContext string) *AppError {
	return &AppError{
		Err:         err,
		Message:     message,
		Code:        code,
		Details:     details,
		Timestamp:   timestamp,
		UserContext: userContext,
	}
}

// NewUnauthorizedError 创建未经授权的应用错误
func NewUnauthorizedError(err error, message string, code int, details, timestamp, userContext string) *AppError {
	return &AppError{
		Err:         err,
		Message:     message,
		Code:        code,
		Details:     details,
		Timestamp:   timestamp,
		UserContext: userContext,
	}
}

// NewTimeoutError 创建超时相关的应用错误
func NewTimeoutError(err error, message string, code int, details, timestamp, userContext string) *AppError {
	return &AppError{
		Err:         err,
		Message:     message,
		Code:        code,
		Details:     details,
		Timestamp:   timestamp,
		UserContext: userContext,
	}
}

// IsAppError 检查错误是否是 AppError 类型
func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

// IsNetworkError 检查错误是否是网络相关的错误
func IsNetworkError(err error) bool {
	appErr, ok := err.(*AppError)
	return ok && appErr.Code == 1002
}

// IsDatabaseError 检查错误是否是数据库相关的错误
func IsDatabaseError(err error) bool {
	appErr, ok := err.(*AppError)
	return ok && appErr.Code == 1003
}

// IsTimeoutError 检查错误是否是超时错误
func IsTimeoutError(err error) bool {
	appErr, ok := err.(*AppError)
	return ok && appErr.Code == 1004
}

// IsUnauthorizedError 检查错误是否是未经授权的错误
func IsUnauthorizedError(err error) bool {
	appErr, ok := err.(*AppError)
	return ok && appErr.Code == 1005
}

// LogError 打印错误日志
func LogError(err error) {
	if appErr, ok := err.(*AppError); ok {
		fmt.Printf("Error [Code: %d, Timestamp: %s, User: %s]: %v\n", appErr.Code, appErr.Timestamp, appErr.UserContext, appErr)
	} else {
		fmt.Println("Error:", err)
	}
}

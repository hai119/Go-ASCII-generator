package errors

import "fmt"

// Error types
var (
    ErrInvalidInput     = fmt.Errorf("invalid input")
    ErrInvalidOutput    = fmt.Errorf("invalid output")
    ErrUnsupportedMode  = fmt.Errorf("unsupported mode")
    ErrFontNotFound     = fmt.Errorf("font not found")
    ErrProcessing       = fmt.Errorf("processing error")
)

// AppError 应用错误结构
type AppError struct {
    Err     error
    Message string
    Code    int
}

func (e *AppError) Error() string {
    return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

// NewAppError 创建新的应用错误
func NewAppError(err error, message string, code int) *AppError {
    return &AppError{
        Err:     err,
        Message: message,
        Code:    code,
    }
}

// WrapError 包装错误
func WrapError(err error, message string) error {
    if err == nil {
        return nil
    }
    return fmt.Errorf("%s: %w", message, err)
} 
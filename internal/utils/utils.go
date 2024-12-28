package utils

import (
    "fmt"
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "path/filepath"
    "strings"
)

// IsSupportedImageFormat 检查是否支持的图像格式
func IsSupportedImageFormat(filename string) bool {
    ext := strings.ToLower(filepath.Ext(filename))
    return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

// IsSupportedVideoFormat 检查是否支持的视频格式
func IsSupportedVideoFormat(filename string) bool {
    ext := strings.ToLower(filepath.Ext(filename))
    return ext == ".mp4" || ext == ".avi" || ext == ".mov"
}

// SaveImage 保存图像
func SaveImage(img image.Image, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer file.Close()

    ext := strings.ToLower(filepath.Ext(filename))
    switch ext {
    case ".jpg", ".jpeg":
        return jpeg.Encode(file, img, nil)
    case ".png":
        return png.Encode(file, img)
    default:
        return fmt.Errorf("unsupported image format: %s", ext)
    }
}

// EnsureDir ��保目录存在
func EnsureDir(path string) error {
    return os.MkdirAll(path, 0755)
}

// GetFileSize 获取文件大小
func GetFileSize(filename string) (int64, error) {
    info, err := os.Stat(filename)
    if err != nil {
        return 0, err
    }
    return info.Size(), nil
} 
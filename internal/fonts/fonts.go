package fonts

import (
    "fmt"
    "path/filepath"

    "github.com/fogleman/gg"
)

// FontConfig 字体配置结构
type FontConfig struct {
    Path string
    Size float64
}

// GetFontConfig 根据语言获取字体配置
func GetFontConfig(language string, scale float64) FontConfig {
    baseSize := 10.0
    switch language {
    case "chinese":
        return FontConfig{
            Path: filepath.Join("fonts", "simsun.ttc"),
            Size: baseSize * scale,
        }
    case "japanese", "korean":
        return FontConfig{
            Path: filepath.Join("fonts", "arial-unicode.ttf"),
            Size: baseSize * scale,
        }
    default:
        return FontConfig{
            Path: filepath.Join("fonts", "DejaVuSansMono-Bold.ttf"),
            Size: baseSize * 2 * scale,
        }
    }
}

// LoadFont 加载字体
func LoadFont(ctx *gg.Context, cfg FontConfig) error {
    if err := ctx.LoadFontFace(cfg.Path, cfg.Size); err != nil {
        return fmt.Errorf("failed to load font %s: %v", cfg.Path, err)
    }
    return nil
} 
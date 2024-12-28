package fonts

import (
    "fmt"
    "path/filepath"
    "sync"

    "github.com/fogleman/gg"
    "github.com/hai119/Go-ASCII-generator/internal/config"
)

// FontManager 字体管理器
type FontManager struct {
    config *config.AppConfig
    cache  map[string]*gg.Context
    mu     sync.RWMutex
}

// NewFontManager 创建字体管理器
func NewFontManager(cfg *config.AppConfig) *FontManager {
    return &FontManager{
        config: cfg,
        cache:  make(map[string]*gg.Context),
    }
}

// GetFont 获取字体上下文
func (m *FontManager) GetFont(language string, size float64) (*gg.Context, error) {
    key := fmt.Sprintf("%s_%f", language, size)

    m.mu.RLock()
    if ctx, ok := m.cache[key]; ok {
        m.mu.RUnlock()
        return ctx, nil
    }
    m.mu.RUnlock()

    m.mu.Lock()
    defer m.mu.Unlock()

    // 再次检查，避免并发创建
    if ctx, ok := m.cache[key]; ok {
        return ctx, nil
    }

    // 创建新的字体上下文
    ctx := gg.NewContext(1, 1) // 临时上下文
    fontPath := m.getFontPath(language)
    if err := ctx.LoadFontFace(fontPath, size); err != nil {
        return nil, fmt.Errorf("failed to load font %s: %w", fontPath, err)
    }

    m.cache[key] = ctx
    return ctx, nil
}

// getFontPath 获取字体路径
func (m *FontManager) getFontPath(language string) string {
    var fontFile string
    switch language {
    case "chinese":
        fontFile = m.config.Fonts.Files["chinese"]
    case "japanese", "korean":
        fontFile = m.config.Fonts.Files["cjk"]
    default:
        fontFile = m.config.Fonts.Files["latin"]
    }
    return filepath.Join(m.config.Fonts.BasePath, fontFile)
}

// ClearCache 清除缓存
func (m *FontManager) ClearCache() {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.cache = make(map[string]*gg.Context)
} 
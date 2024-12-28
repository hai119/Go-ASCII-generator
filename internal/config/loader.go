package config

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v2"
)

// AppConfig 应用配置结构
type AppConfig struct {
    App struct {
        Name    string `yaml:"name"`
        Version string `yaml:"version"`
    } `yaml:"app"`

    Defaults struct {
        Mode          string  `yaml:"mode"`
        Background    string  `yaml:"background"`
        NumCols      int     `yaml:"num_cols"`
        Scale        float64 `yaml:"scale"`
        FPS          int     `yaml:"fps"`
        OverlayRatio float64 `yaml:"overlay_ratio"`
        Language     string  `yaml:"language"`
    } `yaml:"defaults"`

    Fonts struct {
        BasePath    string            `yaml:"base_path"`
        DefaultSize float64           `yaml:"default_size"`
        Files       map[string]string `yaml:"files"`
    } `yaml:"fonts"`

    Output struct {
        SupportedFormats []string `yaml:"supported_formats"`
    } `yaml:"output"`
}

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*AppConfig, error) {
    data, err := os.ReadFile(configPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }

    var config AppConfig
    if err := yaml.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("failed to parse config file: %w", err)
    }

    return &config, nil
}

// MergeWithFlags 将命令行参数与配置文件合并
func MergeWithFlags(cfg *Config, appCfg *AppConfig) *Config {
    if cfg.Mode == "" {
        cfg.Mode = appCfg.Defaults.Mode
    }
    if cfg.Background == "" {
        cfg.Background = appCfg.Defaults.Background
    }
    if cfg.NumCols == 0 {
        cfg.NumCols = appCfg.Defaults.NumCols
    }
    if cfg.Scale == 0 {
        cfg.Scale = appCfg.Defaults.Scale
    }
    if cfg.FPS == 0 {
        cfg.FPS = appCfg.Defaults.FPS
    }
    if cfg.OverlayRatio == 0 {
        cfg.OverlayRatio = appCfg.Defaults.OverlayRatio
    }
    if cfg.Language == "" {
        cfg.Language = appCfg.Defaults.Language
    }
    return cfg
} 
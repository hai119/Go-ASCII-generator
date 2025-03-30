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
		Mode         string  `yaml:"mode"`
		Background   string  `yaml:"background"`
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
	// 检查配置文件是否存在
	if !fileExists(configPath) {
		return nil, fmt.Errorf("config file does not exist: %s", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config AppConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// 验证配置文件中的必填字段
	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
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

	// 如果启用调试模式，输出最终合并的配置
	if isVerboseMode() {
		fmt.Println("Configuration after merging with flags:")
		printConfig(cfg)
	}

	return cfg
}

// fileExists 检查文件是否存在
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// validateConfig 验证配置文件中的必填字段
func validateConfig(config *AppConfig) error {
	if config.App.Name == "" {
		return fmt.Errorf("app name is required")
	}
	if config.App.Version == "" {
		return fmt.Errorf("app version is required")
	}
	if config.Defaults.Mode == "" {
		return fmt.Errorf("default mode is required")
	}
	if config.Fonts.BasePath == "" {
		return fmt.Errorf("fonts base path is required")
	}
	return nil
}

package config

import (
    "flag"
    "os"
    "path/filepath"
)

type Config struct {
    InputPath     string
    OutputPath    string
    Mode          string
    NumCols       int
    Background    string
    CharMode      string
    Scale         float64
    FPS           int
    OverlayRatio  float64
    Language      string
}

// ParseFlags parses command line flags and processes paths
func ParseFlags() *Config {
    cfg := &Config{}
    
    flag.StringVar(&cfg.InputPath, "input", "data/input.jpg", "Path to input file")
    flag.StringVar(&cfg.OutputPath, "output", "data/output.txt", "Path to output file")
    flag.StringVar(&cfg.Mode, "mode", "image2text", "Conversion mode: image2text/image2image/video2video")
    flag.IntVar(&cfg.NumCols, "cols", 100, "Number of columns in output")
    flag.StringVar(&cfg.Background, "bg", "black", "Background color: black/white")
    flag.StringVar(&cfg.CharMode, "char-mode", "complex", "Character set: simple/complex")
    flag.Float64Var(&cfg.Scale, "scale", 1.0, "Output scale")
    flag.IntVar(&cfg.FPS, "fps", 0, "Frames per second (for video)")
    flag.Float64Var(&cfg.OverlayRatio, "overlay", 0.2, "Overlay ratio for video")
    flag.StringVar(&cfg.Language, "lang", "english", "Language for characters")
    
    flag.Parse()

    // 处理路径
    cfg.ProcessPaths()
    
    return cfg
}

// ProcessPaths processes input and output paths
func (cfg *Config) ProcessPaths() {
    // 获取当前工作目录
    workDir, err := os.Getwd()
    if err != nil {
        workDir = "."
    }

    // 处理输入路径
    if !filepath.IsAbs(cfg.InputPath) {
        cfg.InputPath = filepath.Join(workDir, cfg.InputPath)
    }

    // 处理输出路径
    if !filepath.IsAbs(cfg.OutputPath) {
        cfg.OutputPath = filepath.Join(workDir, cfg.OutputPath)
    }

    // 确保输出目录存在
    outputDir := filepath.Dir(cfg.OutputPath)
    if err := os.MkdirAll(outputDir, 0755); err != nil {
        // 如果创建目录失败，记录错误但继续执行
        // 实际写文件时可能会失败
        println("Warning: failed to create output directory:", err)
    }
} 
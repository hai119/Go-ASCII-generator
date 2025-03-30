package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	InputPath    string
	OutputPath   string
	Mode         string
	NumCols      int
	Background   string
	CharMode     string
	Scale        float64
	FPS          int
	OverlayRatio float64
	Language     string
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

	// Print initial debug message for verbosity
	if isVerboseMode() {
		fmt.Println("Flags successfully parsed, starting to process paths...")
	}

	// 处理路径
	cfg.ProcessPaths()

	// After paths are processed, check if mode is valid
	if !isValidMode(cfg.Mode) {
		fmt.Println("Invalid mode specified. Using default mode: image2text")
		cfg.Mode = "image2text"
	}

	// Optionally print the final configuration
	if isVerboseMode() {
		fmt.Println("Configuration processed successfully. Final values:")
		printConfig(cfg)
	}

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
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Printf("Warning: failed to create output directory: %v\n", err)
	}
}

// isVerboseMode checks if verbose mode is enabled
func isVerboseMode() bool {
	return strings.ToLower(os.Getenv("VERBOSE_MODE")) == "true"
}

// isValidMode checks if the specified mode is valid
func isValidMode(mode string) bool {
	validModes := []string{"image2text", "image2image", "video2video"}
	for _, valid := range validModes {
		if mode == valid {
			return true
		}
	}
	return false
}

// printConfig prints the configuration in a readable format
func printConfig(cfg *Config) {
	fmt.Printf("Input Path: %s\n", cfg.InputPath)
	fmt.Printf("Output Path: %s\n", cfg.OutputPath)
	fmt.Printf("Mode: %s\n", cfg.Mode)
	fmt.Printf("Columns: %d\n", cfg.NumCols)
	fmt.Printf("Background: %s\n", cfg.Background)
	fmt.Printf("Character Mode: %s\n", cfg.CharMode)
	fmt.Printf("Scale: %f\n", cfg.Scale)
	fmt.Printf("FPS: %d\n", cfg.FPS)
	fmt.Printf("Overlay Ratio: %f\n", cfg.OverlayRatio)
	fmt.Printf("Language: %s\n", cfg.Language)
}

// RetryOperation attempts an operation multiple times in case of failure
func RetryOperation(operation func() error, retries int, delay time.Duration) error {
	var err error
	for i := 0; i < retries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		fmt.Printf("Attempt %d failed, retrying in %s...\n", i+1, delay)
		time.Sleep(delay)
	}
	return fmt.Errorf("operation failed after %d retries: %w", retries, err)
}

// DebugMode logs the configuration if debug mode is on
func DebugMode() {
	if isVerboseMode() {
		fmt.Println("Debug Mode Active")
	}
}

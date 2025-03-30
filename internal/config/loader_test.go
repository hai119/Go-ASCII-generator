package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	// Create a valid temporary config file for testing
	validConfigYAML := `
app:
  name: MyApp
  version: 1.0
defaults:
  mode: dark
  background: black
  num_cols: 10
  scale: 1.5
  fps: 60
  overlay_ratio: 0.8
  language: en
fonts:
  base_path: /path/to/fonts
  default_size: 12
  files:
    sans: /path/to/sans.ttf
output:
  supported_formats:
    - png
    - jpg
`
	// Temp file path for the valid config
	tmpFile, err := os.CreateTemp("", "valid_config_*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(validConfigYAML)
	require.NoError(t, err)

	// Test loading the valid config file
	t.Run("Valid config", func(t *testing.T) {
		config, err := LoadConfig(tmpFile.Name())
		require.NoError(t, err)
		assert.Equal(t, "MyApp", config.App.Name)
		assert.Equal(t, "1.0", config.App.Version)
		assert.Equal(t, "dark", config.Defaults.Mode)
		assert.Equal(t, 10, config.Defaults.NumCols)
		assert.Equal(t, 1.5, config.Defaults.Scale)
		assert.Equal(t, 60, config.Defaults.FPS)
		assert.Equal(t, 0.8, config.Defaults.OverlayRatio)
		assert.Equal(t, "en", config.Defaults.Language)
		assert.Equal(t, "/path/to/fonts", config.Fonts.BasePath)
		assert.Equal(t, float64(12), config.Fonts.DefaultSize)
		assert.Contains(t, config.Fonts.Files, "sans")
		assert.Equal(t, "/path/to/sans.ttf", config.Fonts.Files["sans"])
		assert.Contains(t, config.Output.SupportedFormats, "png")
		assert.Contains(t, config.Output.SupportedFormats, "jpg")
	})

	// Test loading a non-existent file
	t.Run("File does not exist", func(t *testing.T) {
		_, err := LoadConfig("non_existent_file.yaml")
		assert.Error(t, err)
	})

	// Test loading an invalid config (missing app name)
	invalidConfigYAML := `
app:
  version: 1.0
defaults:
  mode: dark
  background: black
  num_cols: 10
  scale: 1.5
  fps: 60
  overlay_ratio: 0.8
  language: en
fonts:
  base_path: /path/to/fonts
  default_size: 12
  files:
    sans: /path/to/sans.ttf
output:
  supported_formats:
    - png
    - jpg
`
	// Temp file path for the invalid config
	tmpFileInvalid, err := os.CreateTemp("", "invalid_config_*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFileInvalid.Name())

	_, err = tmpFileInvalid.WriteString(invalidConfigYAML)
	require.NoError(t, err)

	t.Run("Invalid config", func(t *testing.T) {
		_, err := LoadConfig(tmpFileInvalid.Name())
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "app name is required")
	})
}

func TestMergeWithFlags(t *testing.T) {
	// Prepare a sample AppConfig loaded from a valid config file
	validConfig := &AppConfig{
		App: struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		}{
			Name:    "MyApp",
			Version: "1.0",
		},
		Defaults: struct {
			Mode         string  `yaml:"mode"`
			Background   string  `yaml:"background"`
			NumCols      int     `yaml:"num_cols"`
			Scale        float64 `yaml:"scale"`
			FPS          int     `yaml:"fps"`
			OverlayRatio float64 `yaml:"overlay_ratio"`
			Language     string  `yaml:"language"`
		}{
			Mode:         "dark",
			Background:   "black",
			NumCols:      10,
			Scale:        1.5,
			FPS:          60,
			OverlayRatio: 0.8,
			Language:     "en",
		},
		Fonts: struct {
			BasePath    string            `yaml:"base_path"`
			DefaultSize float64           `yaml:"default_size"`
			Files       map[string]string `yaml:"files"`
		}{
			BasePath:    "/path/to/fonts",
			DefaultSize: 12,
			Files: map[string]string{
				"sans": "/path/to/sans.ttf",
			},
		},
		Output: struct {
			SupportedFormats []string `yaml:"supported_formats"`
		}{
			SupportedFormats: []string{"png", "jpg"},
		},
	}

	// Test merging flags with config
	t.Run("Merge with flags", func(t *testing.T) {
		cfg := &Config{
			Mode:       "",     // Empty value to fall back on default
			Background: "blue", // Custom value
			NumCols:    0,      // Empty value to fall back on default
		}

		// Merge with the valid config
		mergedConfig := MergeWithFlags(cfg, validConfig)

		assert.Equal(t, "dark", mergedConfig.Mode)       // Should fallback to default
		assert.Equal(t, "blue", mergedConfig.Background) // Should be overridden by the flag
		assert.Equal(t, 10, mergedConfig.NumCols)        // Should fallback to default
	})

	// Test merging with verbose mode enabled
	t.Run("Merge with verbose mode", func(t *testing.T) {
		// Enable verbose mode by setting an environment variable (or mock a function in tests)
		t.Setenv("VERBOSE", "true")

		cfg := &Config{
			Mode: "",
		}

		// Merge with the valid config
		mergedConfig := MergeWithFlags(cfg, validConfig)

		// Here you would validate the printed output, but it's tricky to assert stdout in Go tests.
		// Alternatively, you could mock the printConfig function or test for the side effects.
		assert.Equal(t, "dark", mergedConfig.Mode)
	})

	// Test merging with all fields populated
	t.Run("All fields populated", func(t *testing.T) {
		cfg := &Config{
			Mode:         "light",
			Background:   "white",
			NumCols:      20,
			Scale:        2.0,
			FPS:          120,
			OverlayRatio: 1.0,
			Language:     "fr",
		}

		mergedConfig := MergeWithFlags(cfg, validConfig)

		assert.Equal(t, "light", mergedConfig.Mode)
		assert.Equal(t, "white", mergedConfig.Background)
		assert.Equal(t, 20, mergedConfig.NumCols)
		assert.Equal(t, 2.0, mergedConfig.Scale)
		assert.Equal(t, 120, mergedConfig.FPS)
		assert.Equal(t, 1.0, mergedConfig.OverlayRatio)
		assert.Equal(t, "fr", mergedConfig.Language)
	})
}

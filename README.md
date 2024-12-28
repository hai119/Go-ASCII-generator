# ASCII Art Generator | ASCII 艺术生成器

[English](#english) | [中文](#chinese)

<a name="english"></a>
## English Version

### Introduction

This is a high-performance Go implementation of ASCII art generator, inspired by [Viet Nguyen's Python version](https://github.com/vietnguyen91/ASCII-generator.git). The program utilizes Go's concurrent processing capabilities to provide efficient image and video conversion.

### Key Features

- **Multiple Conversion Modes**
  - Image to ASCII text (.txt)
  - Image to colored ASCII art (.jpg, .png)
  - Video to ASCII text
  - Video to colored ASCII video
  - Support for relative and absolute paths

- **Performance Optimizations**
  - Parallel processing with worker pools
  - Memory optimization
  - Font caching
  - Efficient image processing

- **Customization Options**
  - Multiple languages support (English, Chinese, Japanese, etc.)
  - Background color selection (black/white)
  - Character set options (simple/complex)
  - Output scale adjustment
  - Video frame rate control
  - Video overlay ratio

### System Requirements

- Go 1.21 or higher
- FFmpeg (for video processing)
- Make (for build process)

### Installation

Ubuntu/Debian:
```bash
# Install Go
sudo apt-get update
sudo apt-get install golang-go

# Install FFmpeg
sudo apt-get install ffmpeg
```

macOS:
```bash
# Using Homebrew
brew install go ffmpeg
```

### Building from Source

```bash
# Clone repository
git clone https://github.com/hai119/Go-ASCII-generator.git
cd Go-ASCII-generator

# Build
make build
```

### Usage Examples

1. Basic Image Conversion:
```bash
# Convert image to ASCII text
./bin/ascii --mode image2text --input examples/input.jpg --output output.txt

# Convert image to colored ASCII art
./bin/ascii --mode image2image --input examples/input.jpg --output output.jpg \
        --cols 150 --bg white --char-mode complex
```

2. Video Processing:
```bash
# Convert video to ASCII text
./bin/ascii --mode video2text --input examples/input.mp4 --output output.txt \
        --cols 80 --fps 30

# Convert video to colored ASCII video
./bin/ascii --mode video2video --input examples/input.mp4 --output output.mp4 \
        --cols 100 --scale 1.5 --overlay 0.2
```

### Command Line Options

| Option | Description | Default | Example Values |
|--------|-------------|---------|----------------|
| --mode | Conversion mode | image2text | image2text, image2image, video2text, video2video |
| --input | Input file path | data/input.jpg | Any valid file path |
| --output | Output file path | data/output.txt | Any valid file path |
| --cols | Number of columns | 100 | 80-200 recommended |
| --bg | Background color | black | black, white |
| --char-mode | Character set | complex | simple, complex |
| --scale | Output scale | 1.0 | 0.5-2.0 recommended |
| --fps | Video frame rate | 10 | 1-60 |
| --overlay | Video overlay ratio | 0.2 | 0.0-1.0 |
| --lang | Character set language | english | english, chinese, japanese |

### Project Structure
```
.
├── cmd/
│   └── ascii/          # Main program entry
├── internal/
│   ├── config/         # Configuration management
│   │   └── config.go   # Command line parsing and config
│   └── converter/      # Core conversion logic
│       ├── converter.go    # Image processing
│       ├── image_color.go  # Colored image processing
│       ├── video.go        # Video processing
│       ├── video_color.go  # Colored video processing
│       ├── worker.go       # Worker pool implementation
│       └── utils.go        # Utility functions
└── examples/           # Example files
```

---

<a name="chinese"></a>
## 中文版本

### 简介

这是一个使用 Go 语言实现的高性能 ASCII 艺术生成器，灵感来自 [Viet Nguyen 的 Python 版本](https://github.com/vietnguyen91/ASCII-generator.git)。程序充分利用 Go 的并发处理能力，提供高效的图像和视频转换功能。

### 主要特性

- **多种转换模式**
  - 图片转 ASCII 文本（.txt）
  - 图片转彩色 ASCII 艺术图（.jpg, .png）
  - 视频转 ASCII 文本
  - 视频转彩色 ASCII 视频
  - 支持相对路径和绝对路径

- **性能优化**
  - 使用工作池进行并行处理
  - 内存使用优化
  - 字体缓存
  - 高效的图像处理

- **自定义选项**
  - 多语言支持（英文、中文、日文等）
  - 背景颜色选择（黑/白）
  - 字符集选项（简单/复杂）
  - 输出比例调整
  - 视频帧率控制
  - 视频叠加比例

### 系统要求

- Go 1.21 或更高版本
- FFmpeg（用于视频处理）
- Make（用于构建）

### 安装说明

Ubuntu/Debian 系统：
```bash
# 安装 Go
sudo apt-get update
sudo apt-get install golang-go

# 安装 FFmpeg
sudo apt-get install ffmpeg
```

macOS 系统：
```bash
# 使用 Homebrew 安装
brew install go ffmpeg
```

### 从源码构建

```bash
# 克隆仓库
git clone https://github.com/hai119/Go-ASCII-generator.git
cd Go-ASCII-generator

# 构建
make build
```

### 使用示例

1. 基础图像转换：
```bash
# 图片转 ASCII 文本
./bin/ascii --mode image2text --input examples/input.jpg --output output.txt

# 图片转彩色 ASCII 艺术图
./bin/ascii --mode image2image --input examples/input.jpg --output output.jpg \
        --cols 150 --bg white --char-mode complex
```

2. 视频处理：
```bash
# 视频转 ASCII 文本
./bin/ascii --mode video2text --input examples/input.mp4 --output output.txt \
        --cols 80 --fps 30

# 视频转彩色 ASCII 视频
./bin/ascii --mode video2video --input examples/input.mp4 --output output.mp4 \
        --cols 100 --scale 1.5 --overlay 0.2
```

### 命令行选项

| 选项 | 说明 | 默认值 | 示例值 |
|------|------|--------|--------|
| --mode | 转换模式 | image2text | image2text, image2image, video2text, video2video |
| --input | 输入文件路径 | data/input.jpg | 任意有效文件路径 |
| --output | 输出文件路径 | data/output.txt | 任意有效文件路径 |
| --cols | 输出列数 | 100 | 推荐 80-200 |
| --bg | 背景颜色 | black | black, white |
| --char-mode | 字符集 | complex | simple, complex |
| --scale | 输出比例 | 1.0 | 推荐 0.5-2.0 |
| --fps | 视频帧率 | 10 | 1-60 |
| --overlay | 视频叠加比例 | 0.2 | 0.0-1.0 |
| --lang | 字符集语言 | english | english, chinese, japanese |

### 项目结构
```
.
├── cmd/
│   └── ascii/          # 程序入口
├── internal/
│   ├── config/         # 配置管理
│   │   └── config.go   # 命令行解析和配置
│   └── converter/      # 核心转换逻辑
│       ├── converter.go    # 图像处理
│       ├── image_color.go  # 彩色图像处理
│       ├── video.go        # 视频处理
│       ├── video_color.go  # 彩色视频处理
│       ├── worker.go       # 工作池实现
│       └── utils.go        # 工具函数
└── examples/           # 示例文件
```

### 常见问题

1. Q: 如何提高视频处理速度？  
   A: 可以通过降低 `--cols` 值或降低帧率来提高处理速度。

2. Q: 支持哪些视频格式？  
   A: 通过 FFmpeg 支持常见格式如 MP4、AVI、MOV 等。

3. Q: 如何处理大分辨率图片？  
   A: 可以通过调整 `--scale` 参数来控制输出大小。

### 许可证

MIT License

### 致谢

- 原始 Python 版本作者：Viet Nguyen
- FFmpeg 团队
- Go 图像处理社区

### 联系方式

- 项目维护者：[ZhaoYang Li]
- 项目链接：[https://github.com/hai119/Go-ASCII-generator](https://github.com/hai119/Go-ASCII-generator)

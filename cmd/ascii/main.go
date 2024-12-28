package main

import (
    "fmt"
    "log"

    "github.com/hai119/Go-ASCII-generator/internal/config"
    "github.com/hai119/Go-ASCII-generator/internal/converter"
)

func main() {
    // 解析命令行参数
    cfg := config.ParseFlags()

    // 根据模式选择转换方法
    var err error
    switch cfg.Mode {
    case "image2text":
        err = converter.ImageToText(cfg)
    case "image2image":
        err = converter.ImageToImageColor(cfg)
    case "video2text":
        err = converter.VideoToText(cfg)
    case "video2video":
        err = converter.VideoToVideoColor(cfg)
    default:
        err = fmt.Errorf("unsupported mode: %s", cfg.Mode)
    }

    if err != nil {
        log.Fatal(err)
    }
} 
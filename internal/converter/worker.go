package converter

import (
    "image"
    "image/color"
    "sync"
)

// WorkerPool 工作池结构
type WorkerPool struct {
    numWorkers int
    jobs       chan Job
    results    chan Result
    wg         sync.WaitGroup
}

// Job 工作单元
type Job struct {
    img        image.Image
    row        int
    col        int
    cellWidth  int
    cellHeight int
}

// Result 处理结果
type Result struct {
    row        int
    col        int
    brightness float64
    avgColor   color.Color
}

// NewWorkerPool 创建新的工作池
func NewWorkerPool(numWorkers int) *WorkerPool {
    return &WorkerPool{
        numWorkers: numWorkers,
        jobs:       make(chan Job, numWorkers*2),
        results:    make(chan Result, numWorkers*2),
    }
}

// Start 启动工作池
func (p *WorkerPool) Start() {
    for i := 0; i < p.numWorkers; i++ {
        p.wg.Add(1)
        go p.worker()
    }
}

// Stop 停止工作池
func (p *WorkerPool) Stop() {
    close(p.jobs)
    p.wg.Wait()
    close(p.results)
}

// worker 工作协程
func (p *WorkerPool) worker() {
    defer p.wg.Done()
    for job := range p.jobs {
        x := job.col * job.cellWidth
        y := job.row * job.cellHeight

        brightness := calculateBrightness(job.img,
            x,
            y,
            job.cellWidth,
            job.cellHeight)

        avgColor := calculateAverageColor(job.img,
            x,
            y,
            job.cellWidth,
            job.cellHeight)

        p.results <- Result{
            row:        job.row,
            col:        job.col,
            brightness: brightness,
            avgColor:   avgColor,
        }
    }
} 
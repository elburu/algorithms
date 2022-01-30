package main

import (
	"fmt"
	"math"
	"sync"
)

type JobRunner interface {
	Run()
}

type Job struct {
	fn JobRunner
}

type WorkerPool struct {
	wg    sync.WaitGroup
	queue chan Job
}

func NewWorkerPool(workerCount int) *WorkerPool {
	wp := &WorkerPool{
		queue: make(chan Job),
	}
	wp.addWorkers(workerCount)

	return wp
}

func (wp *WorkerPool) Close() {
	close(wp.queue)
	wp.wg.Wait()
}

func (wp *WorkerPool) ScheduleJob(job JobRunner) {
	wp.queue <- Job{fn: job}
}

func (wp *WorkerPool) addWorkers(workerCount int) {
	wp.wg.Add(workerCount)
	for i := 1; i <= workerCount; i++ {
		go func(jobID int) {
			count := 0
			for job := range wp.queue {
				job.fn.Run()
				count++
			}
			fmt.Println(fmt.Sprintf("Worker#%v -> %v elements processed", jobID, count))

			wp.wg.Done()
		}(i)
	}
}

type Task struct {
	name string
	days int
}

func (t *Task) Run() {
	weeks := t.days / 7
	remDays := math.Remainder(float64(t.days), 7)

	result := fmt.Sprintf("%s ha trabajando en la empresa %v semanas", t.name, weeks)

	if remDays > 0 {
		d := "días"
		if remDays == 1 {
			d = "día"
		}
		result = fmt.Sprintf("%s y %v %s", result, remDays, d)
	}

	fmt.Println(result)
}

func main() {
	numOfWorkers := 1
	e := map[string]int{"Robert": 30, "John": 475, "Elly": 1022, "Bob": 99}

	workerPool := NewWorkerPool(numOfWorkers)

	for name, days := range e {
		workerPool.ScheduleJob(&Task{
			name: name,
			days: days,
		})
	}

	workerPool.Close()
}

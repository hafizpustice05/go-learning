package main

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task interface {
	process()
}
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

func (t EmailTask) process() {
	fmt.Printf("Sendding Email to %s\n", t.Email)

	time.Sleep(time.Second * 2)
}

type ImageProcessingTask struct {
	ImageUrl string
}

func (t *ImageProcessingTask) process() {
	fmt.Printf("Image Processing task %s \n", t.ImageUrl)
	//time consuming task
	time.Sleep(time.Second * 5)
}

// Worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// Function to execute worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.process()
		wp.wg.Done()
	}
}
func (wp *WorkerPool) Run() {
	// Initilize the task chanel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	//send tasks to the taskChanel
	wp.wg.Add(len(wp.Tasks))

	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	//wait for all task to finish
	wp.wg.Wait()
}

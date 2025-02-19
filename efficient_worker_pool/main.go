package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	//create tasks

	tasks := []Task{
		&EmailTask{Email: "email1@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample1.jpg"},
		&EmailTask{Email: "email2@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample2.jpg"},
		&EmailTask{Email: "email3@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample3.jpg"},
		&EmailTask{Email: "email4@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample4.jpg"},
		&EmailTask{Email: "email5@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample5.jpg"},
	}

	// Create a worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, // Number of workers that can run at a time
	}

	wp.Run()
	fmt.Println("Finished the all task")

	fmt.Println("execution time: ", time.Since(now))
}

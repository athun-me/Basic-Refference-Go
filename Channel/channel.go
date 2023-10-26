package main

import "fmt"

func main() {
	numWorkers := 4
	taskQueue := make(chan int)

	for i := 0; i < numWorkers; i++ {
		go worker(i, taskQueue)
	}

	// Add tasks to the queue
	for i := 0; i < 10; i++ {
		taskQueue <- i
	}

	close(taskQueue) // Close the task queue when all tasks are added

	// Wait for all workers to complete
	for i := 0; i < numWorkers; i++ {
		<-taskQueue
	}
}

func worker(id int, taskQueue chan int) {
	for task := range taskQueue {
		fmt.Printf("Worker %d is processing task %d\n", id, task)
		// Perform some work here
		// ...
		taskQueue <- task // Signal task completion
	}
}

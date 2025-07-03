package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup

	// Task Example
	tasks := []task{"Task1", "Task2", "Task3"}

	// Create a channel
	ch := make(chan string, len(tasks))

	// Start a goroutine
	for _, t := range tasks {
		wg.Add(1)
		go func(t task) {
			defer wg.Done()
			// Process the task and send the result to the channel
			result, err := processTask(t)
			if err != nil {
				fmt.Println("Error processing task:", err)
				return
			}
			ch <- result
		}(t)
	}

	// Close the channel after all goroutines finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive the message from the channel
	for msg := range ch {
		println(msg)
	}
}

type task string

func processTask(t task) (string, error) {
	fmt.Println("Processing task:", t)
	time.Sleep(2 * time.Second) // Simulate some processing time
	fmt.Println("Task processed:", t)
	return fmt.Sprintf("Result of %s", t), nil
}
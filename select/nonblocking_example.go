package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Create a buffered channel with capacity for only 2 jobs
	jobsChannel := make(chan string, 2)

	// A dummy list of 5 jobs we want to process
	incomingJobs := []string{
		"Job A",
		"Job B",
		"Job C", // This should fill the buffer
		"Job D", // This should be dropped
		"Job E", // This should be dropped
	}

	for _, job := range incomingJobs {
		// 2. Try to send the job
		select {
		case jobsChannel <- job:
			fmt.Printf("✅ Success: %s added to queue\n", job)
		default:
			// 3. This block runs immediately if the send above would block
			fmt.Printf("❌ Failed: Queue is full! Dropping %s\n", job)
		}
	}

	fmt.Println("\n--- Processing Jobs in Queue ---")

	// Close channel to stop range loop once empty
	close(jobsChannel)

	// Simulate a worker processing the successful jobs
	for job := range jobsChannel {
		fmt.Printf("Working on: %s\n", job)
		time.Sleep(500 * time.Millisecond)
	}
}

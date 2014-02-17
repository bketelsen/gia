package main

import (
	"fmt"
	"time"
)

func main() {
	// Create an unbuffered channel
	baton := make(chan int)

	// First runner to his mark
	go Runner(baton)

	// Start the race
	baton <- 1

	// Give the runners time to race
	time.Sleep(500 * time.Millisecond)
}

func Runner(baton chan int) {
	// Wait to receive the baton
	runner := <-baton

	fmt.Printf("Runner %d Running With Baton\\n", runner)
	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\\n", runner)
		return
	}

	// New runner to the line
	newRunner := runner + 1
	go Runner(baton)

	fmt.Printf("Runner %d Exchange With Runner %d\\n", runner, newRunner)
	baton <- newRunner
}

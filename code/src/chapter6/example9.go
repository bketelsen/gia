package main

import (
	"fmt"
	"time"
)

func main() {
    // Create an unbuffered channel
	ball := make(chan int)

	// Launch two players
	go Player("A", ball)
	go Player("B", ball)

	// Start the lobby
	ball <- 0

	// Give the players time to play
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("Hits: %d\\n", <-ball)
}

func Player(name string, ball chan int) {
	for {
	    // Wait for the ball to be hit back
	    // to us
		value := <-ball

		// Increment the hit count by one
		value++
		fmt.Printf("Player %s Hit %d\\n", name, value)

		// Hit the ball back to the
		// opposing player
		ball <- value
	}
}

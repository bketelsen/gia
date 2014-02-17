package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var Shutdown int64 = 0

func main() {
	// Launch two functions as a goroutine
	go DoWork("A")
	go DoWork("B")

    // Give the goroutines time to run
	time.Sleep(1 * time.Second)

	// Safely flag it is time to shutdown
	atomic.StoreInt64(&Shutdown, 1)

	// Give the goroutines time to shutdown
	time.Sleep(1 * time.Second)
}

func DoWork(name string) {
	for {
		fmt.Printf("Doing %s Work\\n", name)
		time.Sleep(250 * time.Millisecond)

		// Do we need to shutdown
		if atomic.LoadInt64(&Shutdown) == 1 {
			fmt.Printf("Shutting %s Down\\n", name)
			break
		}
	}
}
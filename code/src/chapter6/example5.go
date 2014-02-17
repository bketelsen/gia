package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"runtime"
)

var Counter int64 = 0

func main() {
    // Launch two functions as a goroutine
	go IncCounter(1)
	go IncCounter(2)

    // Give the goroutines time to run
	time.Sleep(1 * time.Second)

	// Display the final value
	fmt.Printf("Final Counter: %d\\n", Counter)
}

func IncCounter(id int) {
	for count := 0; count < 2; count++ {
	    // Safely Add One To Counter
		atomic.AddInt64(&Counter, 1)

        // Yield the processor
		runtime.Gosched()
	}
}
package main

import (
	"fmt"
	"time"
	"runtime"
)

var Counter int = 0

func main() {
    // Launch two functions as a goroutine
	go IncCounter(1)
	go IncCounter(2)

    // Give the goroutines time to run
	time.Sleep(1 * time.Second)
	fmt.Printf("Final Counter: %d\\n", Counter)
}

func IncCounter(id int) {
	for count := 0; count < 2; count++ {
		// Capture the value of Counter
		value := Counter

		// Yield the processor
		runtime.Gosched()

		// Increment our local value of Counter
		value++

		// Store the value back into Counter
		Counter = value
	}
}

package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

var Counter int = 0
var Mutex sync.Mutex  // Create a mutex

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
	    // Only allow one goroutine through this
	    // critical section at a time
	    Mutex.Lock()

		// Capture the value of Counter
		value := Counter

		// Yield the processor
		runtime.Gosched()

		// Increment our local value of Counter
		value++

		// Store the value back into Counter
		Counter = value

		// Release the lock and allow any
		// waiting goroutine through
		Mutex.Unlock()
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting Go Routines")
	
	// Launch two functions as a goroutine
	go PrintPrime("A")
	go PrintPrime("B")

    // Give the goroutines time to run
	fmt.Println("Waiting To Finish")
	time.Sleep(1 * time.Second)
	
	fmt.Println("Terminating Program")
}

func PrintPrime(prefix string) {
next:
	for outer := 1; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\\n", prefix, outer)
	}
	fmt.Printf("Completed %s\\n", prefix)
}

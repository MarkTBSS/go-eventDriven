package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to communicate events
	eventChan := make(chan int)

	// Start a goroutine to handle events
	go func() {
		for {
			event := <-eventChan
			if event%2 == 0 {
				fmt.Println("Event received:", event, " ==> Even Number")
			} else {
				fmt.Println("Event received:", event, " ==> Odd Number")
			}
		}
	}()

	// Send some events
	for i := 1; i <= 5; i++ {
		eventChan <- i
		time.Sleep(1 * time.Second) // Simulate some processing time
	}

	// Wait for a key press before exiting
	fmt.Println("Press any key to exit...")
	var input string
	fmt.Scanln(&input)
}

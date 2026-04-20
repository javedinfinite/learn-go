package main

import (
    "fmt"
    "time"
)

func PrintRateLimitter () {
	fmt.Println("*******************Printing rate limitter*************************")

	// A rate limiter controls how often something is allowed to happen (e.g., API calls per second). 
	// In Go, this is usually done with channels + time, or with a helper from the standard ecosystem.


	// create number of requests we want to check on
	requests := make(chan int, 5)
	for i:= range(5) {
		requests <- i
	}

	close(requests)

	// Create a channel to use it as rate limitter
	rateLimitter := make(chan time.Time, 3)

	// Fill the buffer of the channel with 3 values to be used as burst
	for  range(3) {
		rateLimitter <- time.Now()
	}

	// Write a go routine  which will keep adding value to the rateLimitter channel after every second so that when there is value request will be sent else wait the value to be added by timer
	go func () {
		for t := range time.Tick(time.Second * 1) { // It cannot be stopped, so it can leak resources in long-running programs. use ticker := time.NewTicker(time.Second)
			rateLimitter <- t // after every second timer data will be written to the channel. And this will keep running in background forever but next value will be written to it only when current existing value is ready else it will wait here
		}
	}()


	// Now lets use the ratelimitter by calling multiple requests we ahd already built, you will see first 3 request instant and rest after 1 second.
	for r := range requests {
		<- rateLimitter
		fmt.Println("Sending request", r, time.Now())
	}




}
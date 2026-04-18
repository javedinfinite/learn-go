package main

import (
	"fmt"
	"time"
)

func PrintingSelect () {
	fmt.Println("************Printing select*******************")

	// Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.
	// select is specifically built for channels.
	// select chooses (Whichever channel talks first, I’ll listen to that one):
	// 1. any ready case
	// 2. if multiple are ready → picks one randomly
	// 3. if none are ready → blocks (unless default exists)
	// Note: When one case in a select is chosen, all other cases are ignored for that iteration.

	channel1 := make(chan string)
	channel2 := make(chan string)

	go func (){
		time.Sleep(time.Second * 1)
		channel1 <- "I am don"
	}()

	go func () {
		time.Sleep(time.Second * 2)
		channel2 <- "I am g-one"
	}()

	for range 2 {
		select {
		case val1 := <- channel1: // If a value is available to receive from channel channel1, take it and store it in val1, then execute this case.
			fmt.Println("Channel1 value : ", val1)
		case val2 := <- channel2:
			fmt.Println("Channel2 value : ", val2)

		}
	} 
}
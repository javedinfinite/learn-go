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

	fmt.Println("*********************Timoeout implementation using select and channel*****************")

	channel3 := make(chan string)

	go func () {
		time.Sleep(time.Second * 4)
		channel3 <- "Something in channel 12321"
		fmt.Println("I am the first go routine")
	}()

	select {
	case <- time.After(time.Second * 3): // This is timeout implementation where we are waiting only for 3 seconds to receive response from channel else we will do something
		fmt.Println("Timeout happended, execute something")
	case <- channel3 :
		fmt.Println("This is channel getting read") // This will execute only when channel is ready before timeout timing else timeout will happen 
	}

	fmt.Println("*****************************Timers and Tickers***********************")

	timer := time.NewTimer(time.Second * 2) // Creates a channel and after the time is over, assigns current time value to the channel
	timerValue := <- timer.C // a channel , waiting till timer is over and time value will be assigned to the channel, similar to <-time.After(2 * time.Second)
	fmt.Println("Timer is fired : ", timerValue)

	ticker := time. NewTicker(time.Second * 1) // creates a channel and after every time mentkioned, it sends current time to the channel
	
	i := 0
	for tick := range ticker.C {
		if(i>3) {
			ticker.Stop()
			break
		}
		fmt.Println("Ticker : ", tick)
		i++
	}


}
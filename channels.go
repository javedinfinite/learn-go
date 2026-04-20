package main

import (
	"fmt"
	"time" )

func ping(pings chan<- string, msg string) {
	pings <- msg
}
	
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func PrintChannels1 () {
	// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
	// Note: In go, Everything runs inside goroutines, including main.
	// 
	fmt.Println("*********************Printing channels***********************************")
	

	channel1 := make(chan string) // unbuffered channel created
	channel1 <- "I am don" // blocks until another goroutine is ready to receive, fatal error: all goroutines are asleep - deadlock!


	// Reson for below code never getting executed :
	// The main function calls PrintChannels, and both execute within the same (main) goroutine.
	// When this goroutine tries to send a value into the unbuffered channel using channel1 <- "I am don", it blocks because no other goroutine is ready to receive from that channel.
	// Since the same goroutine is supposed to execute the receive operation (<-channel1) later, but is currently blocked at the send, it never reaches that point.
	// As a result, no progress can be made, and the program ends in a deadlock.

	// Learning: A single goroutine cannot both send and receive on an unbuffered channel unless those operations happen concurrently. 
	// Since this code is sequential, the send blocks forever.

	channel1Value := <- channel1

	fmt.Println("channel1Value : ", channel1Value)

	channel2 := time.After(3 * time.Second) // Another way to create channel with timer, creates a channel, starts a timer, after 3 seconds go sends the current time into the channel2
	fmt.Println("channel2 : ", channel2)


}

func PrintChannels2 () {
	fmt.Println("*******************Printing go routine without deadlock*********************")

	channel1 := make(chan string)

	go func () {
		fmt.Println("I am goroutine, writing to channel1")
		channel1 <- "I am don!"
		channel1 <- "I can do anything"
		channel1 <- "catch me if you can"
		// learning: The sender should close the channel, not the receiver
		// Reason : The sender knows when it’s done sending. The receiver does not know if more data is coming
		// Closing a channel means: “no more values will be sent”
		close(channel1)
		j, more := <- channel1 // In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received. We use this to notify on done when we’ve worked all our jobs.
		fmt.Println("j, more", j, more)
	}()

	// A channel is not like a file or array, you don’t “read all contents.”
	// Each receive operation gets exactly one value, and you need a loop to consume multiple values.
	channel1Value := <- channel1
	fmt.Println("channelValue1 : ", channel1Value)
	
	for value := range channel1 {
		// read all from channel1 but this works only if the channel is closed using close(ch) otherwise errror "fatal error: all goroutines are asleep - deadlock!"
		// learning: The sender should close the channel, not the receiver
		// Reason : The sender knows when it’s done sending. The receiver does not know if more data is coming
		// If you're not using range, you often don’t need to close.
		// Here range works differently on channel i.e. instead of iterating over a fixed collection, it keeps receiving values until the channel is closed
		fmt.Println("Reading channel1 by iterating using range : ", value) 
	}

	// Learning: In Go, synchronization means making goroutines wait for each other at the right moment and 
	// unbuffered channels do this automatically by blocking send and receive until both sides are ready.
	// In our above example, main goroutine waits here :- "channel1Value := <- channel1" until worker goroutine sends data, 
	// This is synchronization automatically. This auto sync happens only for unbuffered channels. Check go_routine.go where we have used waitgroup to synchronize


	fmt.Println("************Printing channel with buffer*******************")
	// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) 
	// ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

	channel2 := make(chan string, 2)

	channel2 <- "I am g-one"
	channel2 <- "I am a robot"
	fmt.Println("Reading from buffered channel:", <-channel2)
	fmt.Println("Reading from buffered channel:", <-channel2)
	// fmt.Println("Reading from buffered channel:", <-channel2) // This will throw error as buffered data is empty and main goroutine is waiting to receive forever but as there is no data to send, it is deadlock

	fmt.Println("************Printing channel direction*******************")
	// Channel Direction (copy pasted from https://gobyexample.com/channel-directions)
	// When using channels as function parameters, you can specify if a channel is meant to only send or receive values. 
	// This specificity increases the type-safety of the program.

	pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
	fmt.Println(<-pongs)

	fmt.Println("************Read sbout select specifically built for channels sync select.go*******************")

}

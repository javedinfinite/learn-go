package main

import "fmt"


func PrintPanic () {
	fmt.Println("*****************Printing Panic*********************")

	// In Go, a panic is a way for the program to stop normal execution immediately when something goes seriously wrong

	// If something happens due to programmer mistakes or unrecoverable states then use panic 
	// panic("Something unrecoverable happened")


	fmt.Println("*****************Printing Defer*********************")

	// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. 
	// defer is often used where e.g. ensure and finally would be used in other languages.
	// defer schedules a function call to run after the surrounding function finishes (even if there’s a panic).

	// defer followed by a function or method call\
	defer fmt.Println("1. I am the one who runs at last even in panic")
	defer func () {
		fmt.Println("2. I am here to run finally even if there is any panic situation")
	} ()
	defer fmt.Println("3. I am the one who runs at last even in panic") // The last defer gets executed first

	fmt.Printf("Hey! I am here to just get executed \n")
	fmt.Printf("Hey! I am here to to run 2 \n")
	fmt.Printf("Hey! I am here to to run 3 \n")

	// panic("Lets test defer working")

	fmt.Println("**********************Printing Recover*******************************")
	// Go makes it possible to recover from a panic, by using the recover built-in function. 
	// A recover can stop a panic from aborting the program and let it continue with execution instead.
	// recover must be called within a deferred function.
	// The return value of recover is the error raised in the call to panic.

	defer func () {
		r := recover()
		if r!=nil {
			fmt.Println("I am here to recover from panic, here is the error caught lets do something with this error : ", r)
		}
	}()

	panic("Creating this panic to get recovered")

	// fmt.Println("testing after recovery") // Unreachable code as after panic, recover code executes


}

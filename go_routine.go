package main

import (
	"fmt"
	"time"
	"sync"
	)

func goRoutine (caller string) {
	for i:= range(3) {
		fmt.Println("Printing from goRoutine :", caller, i)
		time.Sleep(time.Millisecond * 50) // goroutines pause briefly and switch often, so you can clearly see them taking turns and hence second and third are mixed and not sequential
	}
}

func goRoutine2 (caller string, wg *sync.WaitGroup) {
	for i:=range(3) {
		fmt.Println("Printing from goRoutine :", caller, i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() // Internally reduces counter by 1
}

func PrintingGoRoutine1 () {
	// A goroutine is a lightweight thread of execution.
	fmt.Println("*********************Printing Go routine********************************")

	goRoutine("Sync first")

	go goRoutine("Async second") // goroutine will execute concurrently with the calling one.

	go func(caller string) {
			for i:= range(3) {
		fmt.Println("Printing from goRoutine :", caller, i)
		time.Sleep(time.Millisecond * 50)
	}
	}("Async third")

	time.Sleep(time.Second*2) // keep program alive long enough for goroutines to finish becasue if caller ends before threads finishes, threads will be killed
	fmt.Println("go routine caller func is done")

}

func PrintingGoRoutine2 () {
	// A WaitGroup in Go is used to make the main function (or any goroutine) wait until other goroutines finish their work.

	fmt.Println("*************************Printing go routine with waitgroup***********************")

	goRoutine("Sync first without waitgroup")

	var wg sync.WaitGroup // dclaring waitgroup variable

	wg.Add(2) // Starting counter to let system know that we want to wait for 2 go routines before ending this main program, program counter set to 2 and expecting wait.Done() called 2 times.

	go goRoutine2("Async second", &wg) // if a WaitGroup is explicitly passed into functions, it should be done by pointer.

	go func(caller string) {
		for i:=range(3){
			fmt.Println("Printing from goRoutine :", caller, i)
			time.Sleep(time.Millisecond*50)
		}
		wg.Done() // Internally reduces counter by 1
	}("Async third")

	wg.Wait() // blocks main goroutine until counter becomes 0

	fmt.Println("*************************Printing go routine with waitgroup.go***********************")

	// There is another way to wait for multiple go routines to finish without explicitly handling the counter.

	var wgGo sync.WaitGroup

	for i:=1 ; i<3 ; i++ {
		wgGo.Go( 
			func() { 
				goRoutine("This go routine with waitgroup.go")
			})
	}

	wgGo.Wait()

}
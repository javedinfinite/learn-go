package main

import ("fmt"
		"time")

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}


func PrintWorkerPool () {
	fmt.Println("*****************Printing worker pool********************")

	// Copy pasted code from https://gobyexample.com/worker-pools
    // It's not a go feature, it's a concept that how to utilize channel and go routines features to create some worker pool 
    // Here we launched 3 workers waiting/competing on same channel jobs to get data to be processed as per availability on random basis
	
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= numJobs; a++ {
        r:= <-results
		fmt.Println("results : ", r)
    }
	
}


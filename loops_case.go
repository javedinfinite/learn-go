package main

import "fmt"

func goloops_case () {
	fmt.Println("Learning lopps and case in go")

	// loop using for
	var i=1
	for i<2 {
		fmt.Println("for loop with just condition:", i)
		i++
	}

	for j:= 0 ; j<5 ; j++ {
		fmt.Println("Proper for loop:", j)
	}

	for { // infinite loop
		fmt.Println("I can run infinite times if no break")
		break;
	}

	// for with range: do this n times
	for k:= range 5 { // range intialise with 0 till range-1 i.e. range times here it is 5 times
		fmt.Println("k", k)
		if k%2==0 {
			continue
		} else {
			// just to do something, testing else 
		}
		fmt.Println("I am oddn", k)
			
	}

	fmt.Println("Learning switch case:")

	num :=4
	switch num {
	case 1:
		fmt.Println("Switch case",1)
	case 2: 	
		fmt.Println("Switch case", 2)
	case 3,4:
		fmt.Println("Switch case", 3)
	default:
		fmt.Println("Switch default case")
	}

	switch {
	case num>2:
		fmt.Println("I am switch without expression") // note expression can be anything say type or int or bool or etc
	}

}
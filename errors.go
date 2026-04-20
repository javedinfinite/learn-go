package main

import ("fmt"
		"errors")


func divide (num1 float64, num2 float64) ( float64, error) {
	if num2==0 {
		return 0, errors.New("the 2nd number is zero and we can not devide by zero")
	}
	return num1/num2, nil
}

func PrintErrors () {
	fmt.Println("*******************Printing Errors*******************")
	// In Go, error handling is fundamentally different from other languages because errors are treated as normal values, not special exceptions
	// In Go, functions that can fail return two things: the result and an error

	result, err := divide(21,3)

	if err!=nil {
		fmt.Println("This is error do somethng", err)
		return
	}
	fmt.Println("Result is: ", result)

	// Learning: Go intentionally avoids try-catch blocks because they can hide control flow and make code harder to follow. Instead, Go uses Explicit Checking
	// It uses panic(throw), defer(final), recover(handle panic) , check panic_defer_recover.go file for more

	// In a real API, you don't just print errors; you bubble them up to your top-level handler, 
	// which then decides the correct HTTP status code to return to the user.
	// Example:      http.Error(w, "User not found", http.StatusNotFound) return (w is http.ResponseWriter writer came as param)

}
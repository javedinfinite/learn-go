package main

import "fmt"

// Recursion
func getFactoriaOf(num int) int {
	if num==0 {
		return 1 
	} 
	return num*getFactoriaOf(num-1)
}

func PrintFunc () {
	// functions and variadic function(accepting multiple arguments of a type as array using func(fruit ...string)  are easy.
	// functions requires explicit returns and can retuen multiple values
	
	
	// closure: a function that “remembers” variables from its surrounding scope, even after that scope has finished executing.
	// closure is possible in go becasue go allows dollowings:
	// 1. Functions as values (first class functions)  i.e. you can store a function inside a variable f := func () {}
	// 2. Anonymous function : function without name i.e. func (){}()
	// 3. Access to outer variables (lexical scope) x := 10 f:= func () { fmt.Println(x)}
	// Below is a closure example becasue the function “closes over” (captures and encloses) variables from its surrounding scope.

	x := 2
	closureFunc := func() { // anonymous func stored as value
		x++ // accessing variable from outside
		fmt.Println("Printing x from outside of this function", x)
	}
	closureFunc()
	closureFunc()

	// recursion
	fmt.Println(getFactoriaOf(4))
}
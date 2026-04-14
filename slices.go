package main

import ("fmt"
		"slices")

func PrintSlices () {
	// A slice is a dynamic, flexible view over an array. An array declaration without size mentioned is a slice.
	// A slice is called a “view” because it doesn’t store data itself—it just references and exposes a portion of an underlying array.
	fmt.Println("Printing slices:")

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3) 
	slice3 := slice1[1:4] // not a copy, It shares memory with slice1, it is just referencing existing data, any change in slice3 will change slice1 as well.
	slice31 := slice1[:4]
	slice32 := slice1[1:]

	fmt.Println("slice1, slice2, slice3, slice31, slice32", slice1, slice2, slice3, slice31, slice32)

	slice3[2] = 100
	slice4 := make([]int, 3)
	copy(slice4, slice1) // a way to copy existing slice to a new slice to make sure modifying slice4 won't impact slice3
	var slice5 []int
	for i:=range(3) {
		slice5= append(slice5,i) // Go grows slice capacity exponentially (usually doubling) to minimize reallocations and keep append operations efficient.
		fmt.Println("slice5, len, cap", slice5, len(slice5), cap(slice5))
	}

	slice6 := make ([]int, 0, 5) // Providing capacity (3rd arg) in make lets Go allocate memory upfront, avoiding repeated reallocations when the slice grows, it's an optimization and not requirement
	fmt.Println("slice1, slice3, slice4, slice5, slice6", slice1, slice3, slice4, slice5, slice6)


	// slices package uses
	slice7 := []int{4,6,8,0,1}
	slices.Sort(slice7)
	fmt.Println(slice7, slices.Contains(slice7, 2), slices.Equal(slice6, slice7), slices.Index(slice7,8))
	fmt.Println(slices.Insert(slice7, 2, 9), slices.Delete(slice7, 1,3), slice7)

	slice8 := slices.Clone(slice7)
	fmt.Println(slice8)

	slices.Reverse(slice8)
	fmt.Println(slice8)
}
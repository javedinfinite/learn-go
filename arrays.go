package main

import "fmt"

func PrintArrays () {
	fmt.Println("Printing arrays:")

	var arr [7]int
	arr2 := arr
	arr2[2] = 4
	element2 := arr2[2]
	arr3  := [3]int{1,3,5}
	arr4 := [...]int{1,5,7,8} // here comiler will automatically put the array size as per assigned values
	arr5 := [...]int{1, 2, 3, 5:6, 7} // 5:6 meaning, put zero between previous-index and index-5, at index-5 put 6
	fmt.Println("arr, arr2, element2, len(arr), arr3, arr4, arr5 : \n", arr, arr2, element2, len(arr), arr3, arr4, arr5)

	// multi dimensional array

	arr6 := [3][2]int{{1,2},{3,4},{5,6}}
	element := arr6[0][1]
	arr6[0][1] = 4
	fmt.Println("arr6, element",arr6, element)
}
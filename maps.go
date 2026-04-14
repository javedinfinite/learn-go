package main

import ("fmt" 
		"maps"
		"slices")

func PrintMaps () {
	// map is associative data type like hashes, dictionary
	fmt.Println("Printing maps : ")

	map1 := make(map[string]int)
	map2 := map[int]string{1:"one", 2:"two", 3:"three", 4:"four"}

	map1["num1"] = 1
	map1["num2"] = 2
	map1["num3"] = 3
	map1["num4"] = 3

	element1 := map1["num4"]
	element2 := map2[2]
	nonExistingElement, wasPresent := map1["num10"] // returns default zero value of the type
	mapLength := len(map1)


	fmt.Println("map1, element1, nonExistingElement, mapLength, wasPresent : ", map1, element1, nonExistingElement, mapLength, wasPresent)
	fmt.Println("map2, element2:", map2, element2)

	delete(map1, "num4")
	fmt.Println("map1", map1)

	clear(map1)
	fmt.Println("map1", map1)

	map2KeysIterator := maps.Keys(map2) // learning: maps.Keys() returns an iterator, not a slice—so you must range over it or collect it into a slice to see actual keys.
	map2Keys := slices.Collect(map2KeysIterator)

	fmt.Println("map2KeysIterator, map2Keys", map2KeysIterator, map2Keys) // iterator returns keys but in random order, different everytime 

	for num := range(map2KeysIterator) {
		fmt.Println("num:", num)
	}
	



}
package main

import "fmt"

	type Person struct {
		Name string
		Age int
		Height int
		Country string
		Education string
		Weight float64
	}

	// This function belongs to Person, and operates on a pointer to it. i.e. getHeightMulAge is part of struct Person i.e. person1.getHeightMulAge
	// p -> the variable name (like this in other languages)
	// *Person -> the type it’s attached to called "Pointer receivers"
	// (p *Person) and (p Person) both are fine the one with * has the power to update struct value as well  i.e. you can update values of person1
	// Pointer receivers are what make Go structs behave like real-world objects you can update over time, not just temporary values.
	func (p *Person) getHeightMulAge (age int) int { 
		if(age>0){
			p.Age = age
		}
		return p.Height*p.Age
	}

func printStruct () {
	// structs are typed collections of fields. They’re useful for grouping data together to form records.
	fmt.Println("*************************Printing struct*******************************")



	anonymousStruct := struct  {
		name string
		age int
	} {
		"Tiger",
		30,
	} // it's not reusable as you can not have different values by creating different struct type collection out of this

	person1 := Person{Name: "Ra-one", Age: 30, Height: 5} // Here we can choose to not provide value of other fields and it will have zero value of it's type.
	// person2 := Person{"Raone", 30} // In this way of creation of struct requires all field values to be passed
	person2 := &Person{Name: "G-one"}
	person2.Age = 31
	fmt.Println("person1 : ",person1)
	fmt.Println("person1 height : ",person1.Height)
	fmt.Println("person2 pointer after derefferencing it : ",person2)
	fmt.Printf("pointer to Person struct : %p\n", person2)
	fmt.Println("Age using pointer of Person : ", person2.Age)
	fmt.Println("anonymous struct struct2 : ", anonymousStruct)

	fmt.Println("\n Printing methods : ")
	
	method1Result := person1.getHeightMulAge(0) // method in go i.e. calling method getHeightMulAge on struct person1
	fmt.Println("person1 before age update, method1Result : ", person1, method1Result)
	method2Result := person1.getHeightMulAge(34)
	fmt.Println("person1 after age update, method2Result : ", person1, method2Result)

}


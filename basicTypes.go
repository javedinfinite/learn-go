package main

import (
	"fmt"
	"unsafe"
)

func printGoTypes() {
	fmt.Println("Printing basic types in go:")
	// Go has both var and := because it balances explicit, structured declarations (var) for large systems with concise, fast declarations (:=) for everyday coding inside functions.

	// bool,
	var boolVar bool // default value false
	boolvar2 := 2 > 1
	fmt.Println("Printing boolVar, boolvar2:", boolVar, boolvar2)

	// string,
	var string1 string // default value 1 empty space
	string2 := "I am string2"
	fmt.Println("printing string1, string2:", string1, string2)

	// int (32/64 bits usually) or int8x2x2x2,
	var integer1 int // default value is 0
	integer2 := 12345
	var integer3 int8 = 127 // 127 is within range -128 to 127, these explicit int rarely used, use it when you are memory optimizing precisely i.e. storing so many small values
	var integer4 rune = 5000 // alias for int32
	fmt.Println("Printing integer1, integer2, integer3:", integer1, integer2, integer3, integer4)
	fmt.Printf("%T\n", integer4)

	// uint(32/64 bits usually or uint8x2x2x2)
	var uinteger1 uint
	uninteger2 := uint(20)
	var uinteger3 uint8 = 255 // 255 is within range 0 to 255
	var uinteger4 byte = 255  // alias for uint8
	fmt.Println("printing uinteger1, uninteger2, uinteger3, integer4:", uinteger1, uninteger2, uinteger3, uinteger4)
	fmt.Printf("%T, %T\n", uninteger2, uinteger4) // 	println prints values directly, printf for formatted output

	// uintptr: is a type in go to store raw numeric representation of a memory address, addr := uintptr(ptr)  "it's type conversion from poiter to number i.e. integer type uintptr". ptr value most likely is in hexadecimal value, so uintptr converts hexadecimal to decimal pointer used for arithmetic + system-level operations
	// unsafe.Pointer(ptr) keeps pointer semantics (safe for GC), while uintptr(ptr) turns it into a raw number, which breaks safety guarantees.
	// unsafe.Pointer “I know where the house is, and I’m still connected to it.”, uintptr “I wrote the house address on paper but I forgot it’s a house.”
	// uintptr breaks safety because it removes pointer awareness from Go’s garbage collector, while unsafe.Pointer preserves it—so converting to uintptr too early can lead to invalid memory access.(garbase can remove the actual variable x if it finds No real pointer is referencing x anymore )
	x := 10
	ptr := &x
	p := unsafe.Pointer(ptr) // safe intermediate form
	var uintptr1 uintptr // default 0
	uintptr2 := uintptr(p)
	fmt.Println("Printing uintptr1, ptr, p, uintptr2:", uintptr1, ptr, p,  uintptr2)


	// float32, float64
	// float32 → smaller, faster, less memory
	// float64 → default, precise
	float1 := 10.71234567890
	var float2 float32 = 10.91234567890
	fmt.Println("Printing float1, float2", float1, float2)
	fmt.Printf("%T, %T", float1, float2)

	// complex64 (Uses 32-bit float for real and imaginary parts) complex128 (default, Uses 64-bit float for real and imaginary parts) , complex(5, 7) // 5 + 7i
	complex1 := complex(1.123456789, 2.123456789)
	var complex2 complex64 = complex(1.123456789, 2.123456789)
	fmt.Println("Printing complex1, complex2", complex1, complex2)
	fmt.Printf("%T, %T", complex1, complex2)
}

package main

import "fmt"

func PrintString () {
	fmt.Println("Printing strings:")

	// A string is a read-only slice of bytes

	string1 := "one"
	string2 := "é"

	element1 := string1[0]// Returns byte, not character: Strings are stored as UTF-8 bytes, Hence Indexing a string gives a byte and not the character

	fmt.Println("string1_len, element1 unicode code point in decimal of string1 at 0 :", len(string1), element1) 
	fmt.Println("len of string, string2 unicode code point in decimal of characters at 0 and 1 index",len(string2), string2[0], string2[1]) //  "é" has two char and not 1 so len is 2, this is becasue it is non-ascii character, so indexing is not correct way to work with character and text manipulation

	// So what to do to access the char and not ascii value of it. Lets first understand the threory of characters;
	// 1. Unicode = characters (meaning) : It is a universal character set for all language and symbols. It is the definition of characters, not how they are stored.
											// It assigns each char a point called "Unicode code point" "A" = U+0041 (hex value and in decimal it is 65), "é" = U+00E9  (hex value and in decimal it is 233)
	// 2. ASCII = the "basic English subset" of Unicode (0–127) 
	// 3. UTF-8 = how characters are stored in bytes : UTF-8 is way to encode Unicode in bytes. UTF decides how many bytes to use for the char
					// If the character is ASCII then it will be in 1 byte
					// If the character is non-ascii then it may take 2-4 bytes

	// Question: Now Javed, how all the character stored in computer then?

	// Ans: Actually all the characters are first given a "Unicode code point" by UNICODE :
	// example: "A" = U+0041 (hex value , in decimal it is 65), "é" = U+00E9  (hex value , in decimal it is 233)
	// Now it is evaluated that how many bytes the unicode point will take. For ASCII i.e. 0 to 127 it takes 1 byte, for rest (>127) it takes more than 1 byte
	// So, 233 is then broken into two bytes 11000011 (193) and 10101001 (169) but sum is not 233 becasue it contains some system info.
	// 11000011 (193) and 10101001 (169) => (ignoring markers 110 and 10) 00011 + 10101001 => 00011101001 = 233

	// So, when we access any char of a string by it's index like string1[0], it gives the decimal value of unicode point of the character and not the char itself
	// So, len(string1) is not len of char but len of bytes stored in it.
	// But we need the character to work with it, right?
	// Here comes rune the alias for int32, it stores the decimal value of unicode point and if you print the decimal value using print function, 
	// it will translate/print to actual character.
	// Example: s := "é" , s[0] = 195 s[1] = 169 , r := []rune(s) , 195 169  → decode UTF-8 → U+00E9 → 233 , r[0] = 233, fmt.Printf("%c\n", rune(233)), rune(233) is Unicode code point
	// UTF-8 string → "compressed form", rune → "uncompressed character ID"

	//  converting a UTF-8 string into a slice of Unicode code points, where each code point is stored as an int32.
	r := []rune(string2) // r is a []rune (slice of int32), a dynamic slice of int32 
	fmt.Println("r in raw form : ",r)
	fmt.Printf("\nr value : %c",r)

	// When you range over a string in Go, Go automatically decodes UTF-8, returns runes (Unicode code points), gives byte index where each rune starts
	for i,r:= range string2 {
		fmt.Printf("\r range does rune automatically : %c - %d \n", r, i)
	}

	// Values enclosed in single quotes are rune literals.
	fmt.Println('A')  // 65  one Unicode code point
	fmt.Println("A")  // A   a UTF-8 encoded string (byte sequence)





}
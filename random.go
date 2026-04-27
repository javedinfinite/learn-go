package main

import ("fmt"
		"math/rand/v2"
		"strconv"
		)

func PrintRandom () {

	fmt.Println("********************Printing random number*********************")

	// A computer can’t be truly random on its own. So it uses a formula to pretend to be random.
	// Imagine a machine that shuffles cards: 
	// If you start it with the same seed
	// → it shuffles the cards in the exact same way every time
	// If you change the seed
	// → you get a different shuffle
	// So, rand.Intn(10) , Every time you run the program, you may get the same pattern of numbers.
	// With rand.Seed(time.Now().UnixNano()) rand.Intn(10) Now the starting point changes every time (because time keeps moving), so results feel random.
	// In math/rand/v2, No need to manually create a global seed anymore in simple cases

	fmt.Println("random integer : ",rand.IntN(100)) // 0–99
	fmt.Println("random float : ", rand.Float64()) // 0.0 to 1.0
	fmt.Println(rand.IntN(10) + 1) // in case you don't want 0-10 but 1-10

	// Do we still need seed in v2?
	// For most simple programs → NO manual seeding needed. It automatically uses a good default random source. 
	// But if you want controlled randomness i.e. Same seed → random but same sequence every time, Useful for testing, games, simulations

	// Create a deterministic source. 1, 2 = the starting “position” of the random machine’s internal gears
	src := rand.NewPCG(1,2)  // PCG = a modern pseudo-random generator. (1, 2) = the seed values (two numbers instead of one).
	// Create a random generator from that source
	contRanGen := rand.New(src)
	fmt.Println("controlled rangen : ",contRanGen.IntN(10))
	fmt.Println("controlled rangen : ",contRanGen.IntN(10))

	fmt.Println("*****************Printing number pasrsing*******************")

	parsedInt, _ := strconv.ParseInt("234", 10, 64) // base = 10 (decimal), bit size = 64
	parsedFloat, _ := strconv.ParseFloat("234.786", 64) //  bit size = 64
	parseBinInt, _ := strconv.ParseInt("1010", 2, 64) // binary → decimal
	parseIntSimple, _ := strconv.Atoi("123")
	parsedTostring := strconv.Itoa(123)
	parseFloatString := strconv.FormatFloat(3.14, 'f', 3, 64) // f: normal decimal format, 3 decimal places
	
	fmt.Println("parse int : ", parsedInt)
	fmt.Println("parse float : ", parsedFloat)
	fmt.Println("binary to decimal : ", parseBinInt)
	fmt.Println("Simple parseInt without providing details like base and bit : ", parseIntSimple)
	fmt.Println("int to string pasrsed: ", parsedTostring)
	fmt.Println("parseFloatString: ", parseFloatString)


}
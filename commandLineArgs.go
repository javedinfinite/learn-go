package main

import ("fmt"
		"os"
		"flag")

func PrintCommandLineArgs () {
	fmt.Println("***********************Printing Command LIne Arguments*********************")
	// run <go run . heloCMDargs> or <go run main.go commandLineArgs.go heloCMDargs>

	fmt.Println(os.Args)
	fmt.Println(os.Args[0]) // program name
	// fmt.Println(os.Args[1:]) // the arguments
	// fmt.Println(os.Args[2]) // 2nd  argument

	// Basic flag declarations are available for string, integer, and boolean options. Here we declare a string flag name with a default value "defaultname"
	//  and a short description. This flag.String function returns a string pointer (not a string value);
	namePtr := flag.String("name", "defaultname", "description of name flag")
	numPtr := flag.Int("num", 7, "num desc")
	forkPtr := flag.Bool("fork", false, "fork desc")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	// we need to dereference the pointers with e.g. *wordPtr to get the actual option values.
	// ImportAnt: while running the command with flag values, Put flags BEFORE positional arguments. 
	// Go processes arguments in order, and the flag package stops parsing flags as soon as it hits the first non-flag argument
	// < go run main.go commandLineArgs.go -name=Moon -num=8 -fork=true >
	// Use -h or --help flags to get automatically generated help text for the command-line program. So, <go run main.go commandLineArgs.go -h>
	fmt.Println("namePtr: ", *namePtr)
	fmt.Println("numPtr: ", *numPtr)
	fmt.Println("forkPtr: ", *forkPtr)



}
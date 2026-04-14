package main

func main () {
	// heloWorld()
	// printGoTypes()
	// goloops_case()
	// PrintArrays()
	// PrintSlices()
	// PrintMaps()
	PrintFunc()

}


// Learning-1: Every executable needs a func main() inside a package main. If the package isn't main, the Go compiler doesn't know where to start the program.
// If you want to create library to be used inside executable then don't run it, jut build it <go build file_name>

// Learning-2: A single file cannot have two packages. If you want to write multipe go files in the same folder, you need
// only one main function in one of the file with package main and rest of the files should use same package name main.
// If you want different pakcage name for each file then you need to move each file in a different directory exporting the function (first letter capital)
// And then importing the function to main to run it becasue there is only one entry point.

// Learning-3: If you run just main.go file, go doesn't know about other files in the same package main. You have to let it know that other files exists
// which are part of same package by running all the files like this: <go run .> and you will get error "go: cannot find main module, but found .git/config"
// This error is becasue go needs the project to be intialized as go module “This looks like a project (it has Git), but it’s not initialized as a Go module yet.”
// Run <go mod init learn-go> it will create go.mod file. Now you can run <go run .>
package main

import (
	"fmt"
	"os"
	"strings"
)


func PrintEnvVar () {
	fmt.Println("************************Printing Environment variable*********************")

	os.Setenv("myEnv1", "envValue1")
	os.Setenv("myEnv2", "envValue2")

	fmt.Println("env1:", os.Getenv("myEnv1"))
	fmt.Println("env2:", os.Getenv("myEnv2"))

	for i , envVar := range os.Environ() { // returns all environment as slice of strings
		pair := strings.SplitN(envVar, "=", 2) // This splits the string into 2 parts only  Because values can contain = too: DATABASE_URL=user=admin password=123, First part = key and rest is value
		fmt.Println(" environment var : ", i,  pair)
	}

}
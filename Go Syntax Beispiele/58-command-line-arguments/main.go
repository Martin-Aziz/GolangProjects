package main

import (
	"fmt"
	"os"
)

func main() {
	// get args with program name
	argsWithProg := os.Args

	// get args without program name
	argsWithoutProg := os.Args[1:]

	// get specific arg
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

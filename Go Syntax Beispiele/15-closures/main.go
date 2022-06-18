package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		fmt.Printf("i is: %v\n", &i)
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	nextInt()
	nextInt()
	nextInt()

	newInt := intSeq()
	newInt()
}
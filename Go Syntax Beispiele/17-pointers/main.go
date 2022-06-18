package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval val does not change because we passed in the value of i:", i)

	zeroptr(&i)
	fmt.Println("zeroval val changed because we modified the pointer so now value of i is:", i)

	fmt.Println("pointer address is:", &i)
}
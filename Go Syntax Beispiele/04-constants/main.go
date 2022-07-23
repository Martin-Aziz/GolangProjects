package main

import (
	"fmt"
	"math"
)

const(
	name string ="Jame"
	lastName string ="mss"
) 

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	rl := ratelimit.New(100) // per second

	prev := time.Now()
	for i := 0; i < 1000; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		if now.Sub(prev) > (10 * time.Millisecond) {
			fmt.Println("took too long")
		}
		prev = now
	}

	// Output:
	// 0 0
	// 1 10ms
	// 2 10ms
	// 3 10ms
	// 4 10ms
	// 5 10ms
	// 6 10ms
	// 7 10ms
	// 8 10ms
	// 9 10ms
}
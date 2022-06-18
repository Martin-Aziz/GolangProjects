package main

import (
	"time"
	"fmt"
)

func run(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		fmt.Printf("starting %s\n", d)
		time.Sleep(d)
		c <- 1
	}()
	fmt.Printf("completed %s\n", d)
	return c
}

func main() {
	for i := 0; i < 24; i++ {
		c := run(1 * time.Second)
		<-c
	}
}

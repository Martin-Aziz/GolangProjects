package main

import (
	"fmt"
	"time"

	"https://github.com/azizarash/GolangProjects/tree/main/DigitalClock/screen"
)

func main() {
	screen.Clear()

	for {
		screen.MoveTopLeft()

		w, h := screen.Size()
		fmt.Printf("Width: %d Height: %d\n", w, h)

		fmt.Println(time.Now())

		time.Sleep(time.Second)
	}
}

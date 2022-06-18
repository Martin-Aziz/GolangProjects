package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	// Scan the next token
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		_, err = fmt.Fprintln(os.Stderr, "error:", err)
		check(err)
		os.Exit(1)
	}
}
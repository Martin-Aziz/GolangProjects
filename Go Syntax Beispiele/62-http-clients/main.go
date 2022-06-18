package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// GO standard library has support for http client
func main() {

	resp, err := http.Get("https://jasonganub.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	// print the first 5 lines of the response body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; i < 5; i++ {
		scanner.Scan() // advance to the next token
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

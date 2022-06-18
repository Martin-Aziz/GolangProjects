package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// streamlined error checker
func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	// Load entire file contents
	dat, err := ioutil.ReadFile("51-reading-files/tmp")
	check(err)
	fmt.Println(string(dat))

	// Opening a file allows more control
	f, err := os.Open("51-reading-files/tmp")
	check(err)

	// Read up to 5 bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// Seek to a specific line
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Seek is a way to create your own rewind logic
	_, err = f.Seek(0, 0)
	check(err)

	// Bufio has a buffered reader
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
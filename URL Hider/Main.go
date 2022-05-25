//How to use
//go run Test.go " webpage http://www.mysuperpage.com <-- This is my website!"
package main

import (
	"fmt"
	"os"
)

const (
	link  = "http://"
	nlink = len(link)
	mask  = '*'
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter Text!")
		return
	}

	var (
		text = args[0]
		size = len(text)

		// create a byte buffer directly from the string (text)
		buf = []byte(text)

		in bool
	)

	for i := 0; i < size; i++ {
		// no need to add an artificial http:// prefix
		// it's already there
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			i += nlink
		}

		switch text[i] {
		case ' ', '\t', '\n': // try -> unicode.IsSpace
			in = false
		}

		// when censoring mode: on
		// directly manipulate the bytes on the buffer
		if in {
			buf[i] = mask
		}
	}
	fmt.Println(string(buf))
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Create temp file
	f, err := ioutil.TempFile("", "sample")
	check(err)
	fmt.Println("Temp file name", f.Name())

	// clean up afterwards
	defer os.Remove(f.Name())

	// Write data to the file
	data := []byte{1, 2, 3, 4, 5}
	_, err = f.Write(data)
	check(err)

	// Create a temp dir if there are many temp files
	dirName, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name", dirName)

	defer os.RemoveAll(dirName)

	// Append temp dir name to all temp files
	fname := filepath.Join(dirName, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2, 3, 4}, 0666)
	check(err)

	fmt.Println("Files created in temp dir")
	c, err := ioutil.ReadDir(dirName)
	check(err)

	for _, entry := range c {
		fmt.Println("file: ", entry.Name())
	}
}

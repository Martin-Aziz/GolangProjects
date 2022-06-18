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
	// Create subdir in current workdir
	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	// Helper func to create new files
	createEmptyFile := func(name string) {
		d := []byte("")
		err = ioutil.WriteFile(name, d, 0644)
		check(err)
	}

	createEmptyFile("subdir/file1")

	// Recursively create directories
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	// Create files in those dirs
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// Get files in a dir
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Change dir
	err = os.Chdir("subdir/parent/child")
	check(err)

	// check all contents in current dir
	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Move back to original dir
	err = os.Chdir("../../..")
	check(err)

	// Walk through every file or directory recurisvely and call the visit function to display file info
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
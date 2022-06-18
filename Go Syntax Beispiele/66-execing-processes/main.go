package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// set the binary path
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}

	// set the arguments
	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	// execute the binary with the args
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

package main

import (
	"syscall"
	"os"
	"os/exec"
)

func main() {
	binary, lookErr := exec.LookPath("ls")

	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

/*

// run in *nix environment

*/

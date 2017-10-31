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

// windows 下执行报错
$ go run Basic/execing-processes.go
panic: not supported by windows

goroutine 1 [running]:
main.main()
        D:/Projects/myGitHub/go-practise/go-segments/Basic/execing-processes.go:22 +0xe0
exit status 2

*/

package main

import (
	"syscall"
	"os/exec"
	"os"
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

// 备注在window上运行报错
$ go run basic/execing-processes.go
panic: not supported by windows

goroutine 1 [running]:
main.main()
        D:/Projects/myGitHub/go-practise/go-segments/basic/execing-processes.go:21 +0xe0
exit status 2

*/

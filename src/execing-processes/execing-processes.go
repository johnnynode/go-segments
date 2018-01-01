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
$ go run src/execing-processes/execing-processes.go
total 48
drwxr-xr-x  11 Johnny  staff   352B  1  1 06:23 .
drwxr-xr-x   4 Johnny  staff   128B 11 15 08:23 ..
-rw-r--r--@  1 Johnny  staff   6.0K 11 17 16:16 .DS_Store
-rw-r--r--   1 Johnny  staff   612B 11 11 19:15 .editorconfig
drwxr-xr-x  17 Johnny  staff   544B  1  1 11:16 .git
-rw-r--r--   1 Johnny  staff    60B 11 17 10:17 .gitignore
-rw-r--r--   1 Johnny  staff   1.0K 11 11 19:17 LICENSE
-rw-r--r--   1 Johnny  staff   3.4K  1  1 06:45 README.md
drwxr-xr-x   2 Johnny  staff    64B 11 17 10:16 build
drwxr-xr-x  65 Johnny  staff   2.0K  1  1 11:08 src
drwxr-xr-x   8 Johnny  staff   256B  1  1 06:31 test
*/
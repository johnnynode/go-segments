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

	args := []string {"ls","-a","-l","-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
/*
$ go run Basic/execing-processes.go
total 48
drwxr-xr-x   9 Johnny  staff   288B 11 17 09:01 .
drwxr-xr-x   4 Johnny  staff   128B 11 15 08:23 ..
-rw-r--r--@  1 Johnny  staff   6.0K 11 16 09:38 .DS_Store
-rw-r--r--   1 Johnny  staff   612B 11 11 19:15 .editorconfig
drwxr-xr-x  17 Johnny  staff   544B 11 17 09:24 .git
-rw-r--r--   1 Johnny  staff    47B 11 15 08:26 .gitignore
drwxr-xr-x  69 Johnny  staff   2.2K 11 17 09:21 Basic
-rw-r--r--   1 Johnny  staff   1.0K 11 11 19:17 LICENSE
-rw-r--r--   1 Johnny  staff   2.7K 11 11 19:15 README.md
*/
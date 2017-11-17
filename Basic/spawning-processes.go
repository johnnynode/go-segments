package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbay grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hellp")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

/*
$ go run Basic/spawning-processes.go
> date
2017年11月17日 星期五 09时19分52秒 CST

> grep hellp
hello grep

> ls -a -l -h
total 48
drwxr-xr-x   9 Johnny  staff   288B 11 17 09:01 .
drwxr-xr-x   4 Johnny  staff   128B 11 15 08:23 ..
-rw-r--r--@  1 Johnny  staff   6.0K 11 16 09:38 .DS_Store
-rw-r--r--   1 Johnny  staff   612B 11 11 19:15 .editorconfig
drwxr-xr-x  17 Johnny  staff   544B 11 17 09:19 .git
-rw-r--r--   1 Johnny  staff    47B 11 15 08:26 .gitignore
drwxr-xr-x  68 Johnny  staff   2.1K 11 17 09:15 Basic
-rw-r--r--   1 Johnny  staff   1.0K 11 11 19:17 LICENSE
-rw-r--r--   1 Johnny  staff   2.7K 11 11 19:15 README.md

*/
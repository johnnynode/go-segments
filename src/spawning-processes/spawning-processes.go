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
	grepIn.Write([]byte("hello grep\ngoodby grep"))

	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
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
$ go run src/spawning-processes/spawning-processes.go
> date
2018年 1月 1日 星期一 11时06分06秒 CST

> grep hello
hello grep

> ls -a -l -h
total 48
drwxr-xr-x  11 Johnny  staff   352B  1  1 06:23 .
drwxr-xr-x   4 Johnny  staff   128B 11 15 08:23 ..
-rw-r--r--@  1 Johnny  staff   6.0K 11 17 16:16 .DS_Store
-rw-r--r--   1 Johnny  staff   612B 11 11 19:15 .editorconfig
drwxr-xr-x  17 Johnny  staff   544B  1  1 11:06 .git
-rw-r--r--   1 Johnny  staff    60B 11 17 10:17 .gitignore
-rw-r--r--   1 Johnny  staff   1.0K 11 11 19:17 LICENSE
-rw-r--r--   1 Johnny  staff   3.4K  1  1 06:45 README.md
drwxr-xr-x   2 Johnny  staff    64B 11 17 10:16 build
drwxr-xr-x  64 Johnny  staff   2.0K  1  1 10:59 src
drwxr-xr-x   8 Johnny  staff   256B  1  1 06:31 test
*/
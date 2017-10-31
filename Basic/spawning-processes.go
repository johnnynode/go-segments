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
	grepIn.Write([]byte("hellp grep\ngoodbye grep"))
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
2017年10月31日 周二 21:37:48

> grep hellp

> ls -a -l -h
total 3.3M
drwxr-xr-x 1 Johnny 197121    0 10月 31 21:24 .
drwxr-xr-x 1 Johnny 197121    0 10月 30 11:45 ..
drwxr-xr-x 1 Johnny 197121    0 10月 31 21:37 .git
-rw-r--r-- 1 Johnny 197121   39 10月 27 09:33 .gitignore
drwxr-xr-x 1 Johnny 197121    0 10月 31 21:37 Basic
-rwxr-xr-x 1 Johnny 197121 1.6M 10月 31 21:18 command-line-arguments.exe
-rwxr-xr-x 1 Johnny 197121 1.7M 10月 31 21:24 command-line-flags.exe
-rw-r--r-- 1 Johnny 197121 2.8K 10月 30 11:45 README.md
-rw-r--r-- 1 Johnny 197121    5 10月 31 19:16 test.txt
-rw-r--r-- 1 Johnny 197121    9 10月 31 21:09 writing1.txt
-rw-r--r-- 1 Johnny 197121   21 10月 31 21:09 writing2.tx

*/
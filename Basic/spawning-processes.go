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
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
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

$ go run Basic/spawning-processes.go
> date
2017年11月 6日 周一  6:21:23

> grep hello
hello grep

> ls -a -l -h
total 1.6M
drwxr-xr-x 1 Johnny 197121    0 11月  6 06:05 .
drwxr-xr-x 1 Johnny 197121    0 10月 30 11:45 ..
-rw-r--r-- 1 Johnny 197121  644 11月  3 12:02 .editorconfig
drwxr-xr-x 1 Johnny 197121    0 11月  6 06:21 .git
-rw-r--r-- 1 Johnny 197121   39 10月 27 09:33 .gitignore
drwxr-xr-x 1 Johnny 197121    0 11月  6 06:21 Basic
-rwxr-xr-x 1 Johnny 197121 1.6M 11月  6 06:05 command-line-arguments.exe
-rw-r--r-- 1 Johnny 197121 2.8K 10月 30 11:45 README.md
-rw-r--r-- 1 Johnny 197121    5 11月  4 18:02 test.txt
-rw-r--r-- 1 Johnny 197121    9 11月  6 05:57 test1.txt
-rw-r--r-- 1 Johnny 197121   21 11月  6 05:57 test2.txt

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()

	if (err != nil) {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StderrPipe()
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

$ go run basic/spawning-processes.go
> date
2017年10月27日 周五  9:59:06

> grep hello

> ls -a -l -h
total 33K
drwxr-xr-x 1 Johnny 197121  0 10月 27 09:33 .
drwxr-xr-x 1 Johnny 197121  0 9月  25 08:07 ..
drwxr-xr-x 1 Johnny 197121  0 10月 27 09:59 .git
-rw-r--r-- 1 Johnny 197121 39 10月 27 09:33 .gitignore
drwxr-xr-x 1 Johnny 197121  0 10月 27 09:59 basic


*/


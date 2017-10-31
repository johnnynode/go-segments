package main

import (
	"os"
	"fmt"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

/*

$ go build Basic/command-line-arguments.go

$ ./command-line-arguments.exe a b c d
[D:\Projects\myGitHub\go-practise\go-segments\command-line-arguments.exe a b c d]
[a b c d]
c

*/
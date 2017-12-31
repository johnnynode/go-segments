package main

import "os"
import "fmt"

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

/*
$ go build -o test/command-line-arguments src/command-line-arguments/command-line-arguments.go

$ test/command-line-arguments a b c d
[test/command-line-arguments a b c d]
[a b c d]
c
*/
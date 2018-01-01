package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")
	os.Exit(3)
}

/*
$ go build -o test/exit src/exit/exit.go

$ test/exit

$ echo $?
3
*/
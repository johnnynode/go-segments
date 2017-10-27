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

$ go run basic/exit.go
exit status 3

$ go build basic/exit.go
$ ./exit
$ echo $?
3

*/

package main

import "os"
import "strings"
import "fmt"

func main() {
	os.Setenv("FOO", "1") // set
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	os.Unsetenv("FOO") // remove
	os.Unsetenv("BAR") // remove
	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

/*
$ go run src/environment-variables/environment-variables.go
FOO: 1
BAR:

TERM_PROGRAM
TERM
SHELL
TMPDIR
...

$ BAR=2 go run src/environment-variables/environment-variables.go
FOO: 1
BAR: 2

TERM_PROGRAM
TERM
SHELL
TMPDIR
...

*/

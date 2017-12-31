package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("./test/file")
	if err != nil {
		panic(err)
	}
}

/*
$ go run src/panic/panic.go
panic: a problem

goroutine 1 [running]:
main.main()
        /Users/../go-segments/src/panic/panic.go:6 +0x39
exit status 2
*/
package main

import "os"

func main()  {
	panic("a problem")

	_, err := os.Create("./Basic/panic.txt")
	if err != nil {
		panic(err)
	}
}

/*
$ go run Basic/panic.go
panic: a problem

goroutine 1 [running]:
main.main()
        /Users/Johnny/code/github/go-code/src/go-segments/Basic/panic.go:6 +0x39
exit status 2
*/
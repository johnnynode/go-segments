package main

import "os"

func main() {
	panic("a problem")
	_, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
}

/*

$ go run Basic/panic.go
panic: a problem

goroutine 1 [running]:
main.main()
        /../go-segments/Basic/panic.go:6 +0x6b
exit status 2

*/

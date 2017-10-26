package main

import "os"

func main() {
	panic("a problem")

	_,err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

/*
windows上运行
$ go run basic/panic.go
panic: a problem

goroutine 1 [running]:
main.main()
        D:/Projects/myGitHub/go-practise/go-segments/basic/panic.go:6 +0x6b
exit status 2


*/

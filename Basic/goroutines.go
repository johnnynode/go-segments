package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i ++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input) // 此处等待控制台输入阻塞程序运行包括下面的 goroutine(有时候不会运行)
	go f("hi")
	fmt.Println("done")
}

/*

$ go run Basic/goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
goroutine : 1
going
goroutine : 2
↵
done
hi : 0 // 这里可能不会输出或输出不到3条

*/


package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 关闭非空channel仍可能接受到值
	for elem := range queue {
		fmt.Println(elem)
	}
}

/*

$ go run Basic/range-over-channels.go
one
two

*/

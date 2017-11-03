package main

import "fmt"

func main() {
	messages := make(chan string, 2) // 创建一个channel对象 缓存为2

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*

$ go run Basic/channel-buffering.go
buffered
channel

*/

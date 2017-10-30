package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages) // buffered
	fmt.Println(<-messages) // channel
}

/*

$ go run Basic/channel-buffering.go
buffered
channel

*/
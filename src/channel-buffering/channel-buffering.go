package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
$ go run src/channel-buffering/channel-buffering.go
buffered
channel
*/
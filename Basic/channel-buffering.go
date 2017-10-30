package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages) // buffered
	fmt.Println(<-messages) // channel

	fmt.Println()
	m2 := make(chan string)
	go func () {
		m2 <- "m2-send1"
		m2 <- "m2-send2"
	}()
	fmt.Println(<-m2)
	fmt.Println(<-m2)
}

/*

$ go run Basic/channel-buffering.go
buffered
channel

m2-send1
m2-send2

*/
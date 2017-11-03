package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2) // 睡眠2s
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1): // 接受
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}

/*

$ go run Basic/timeouts.go
timeout 1
result 2

*/
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // 发送
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done // 接收
}

/*

$ go run Basic/channel-synchronization.go
working...done

*/
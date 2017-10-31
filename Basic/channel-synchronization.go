package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done") // 1s 之后输出
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}

/*

$ go run Basic/channel-synchronization.go
working...done

*/

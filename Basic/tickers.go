package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

/*

$ go run Basic/tickers.go
Tick at 2017-11-04 17:07:13.9614476 +0800 CST
Tick at 2017-11-04 17:07:14.4613854 +0800 CST
Tick at 2017-11-04 17:07:14.9619112 +0800 CST
Ticker stopped

*/

package main

import (
	"fmt"
	"time"
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
Tick at 2017-10-31 13:47:43.2290491 +0800 CST
Tick at 2017-10-31 13:47:43.7290372 +0800 CST
Tick at 2017-10-31 13:47:44.2297974 +0800 CST
Ticker stopped

*/
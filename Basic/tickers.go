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
Tick at 2017-11-09 14:57:32.9606713 +0800 CST
Tick at 2017-11-09 14:57:33.4607146 +0800 CST
Tick at 2017-11-09 14:57:33.9617146 +0800 CST
Ticker stopped

*/



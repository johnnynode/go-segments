package main

import "time"
import "fmt"

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
$ go run src/tickers/tickers.go
Tick at 2017-12-31 10:07:24.716928 +0800 CST m=+0.501442331
Tick at 2017-12-31 10:07:25.216746 +0800 CST m=+1.001243293
Tick at 2017-12-31 10:07:25.716433 +0800 CST m=+1.500916828
Ticker stopped
*/
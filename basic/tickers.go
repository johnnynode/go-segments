package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(time.Microsecond * 500)
	go func() {
		for t:= range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}




package main

import (
	"fmt"
	"time"
	"sync/atomic"
)

func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i ++ {
		go func () {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

/*
MacBook-Pro:go-segments Johnny$ go run Basic/atomic-counters.go
ops: 39166

*/
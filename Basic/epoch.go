package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

/*
$ go run Basic/epoch.go
2017-11-15 08:30:21.919602 +0800 CST m=+0.000335831
1510705821
1510705821919
1510705821919602000
2017-11-15 08:30:21 +0800 CST2017-11-15 08:30:21.919602 +0800 CST
*/
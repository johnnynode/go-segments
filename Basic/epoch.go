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
2017-10-31 20:21:14.0441231 +0800 CST
1509452474
1509452474044
1509452474044123100
2017-10-31 20:21:14 +0800 CST
2017-10-31 20:21:14.0441231 +0800 CST

*/
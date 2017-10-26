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

	fmt.Println(time.Unix(secs,0))
	fmt.Println(time.Unix(0, nanos))
}

/*

$ go run basic/epoch.go
2017-10-26 14:42:07.3450407 +0800 CST
1509000127
1509000127345
1509000127345040700
2017-10-26 14:42:07 +0800 CST
2017-10-26 14:42:07.3450407 +0800 CST

*/

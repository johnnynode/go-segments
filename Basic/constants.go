package main

import (
	"math"
	"fmt"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const m = 5E8
	fmt.Println(n == m)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}

/*

$ go run Basic/constants.go
constant
true
6e+11
600000000000
-0.28470407323754404

*/

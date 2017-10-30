package main

import "fmt"

// intSeq 返回一个匿名函数，匿名函数返回int
func intSeq() func() int {
	i := 0
	// return 一个匿名函数
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1
}

/*

$ go run Basic/closures.go
1
2
3
1

*/

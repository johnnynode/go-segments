package main

import "fmt"

func main()  {
	var a string = "initial"
	fmt.Println(a)

	var aa string
	fmt.Println(aa)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}

/*

$ go run Basic/variables.go
initial

1 2
true
0
short

*/
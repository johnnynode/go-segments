package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i ++
	}
	fmt.Println("one------>")

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("two------>")

	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("thr------>")

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println("fou------>")
}

/*
$ go run src/for/for.go
1
2
3
one------>
7
8
9
two------>
loop
thr------>
1
3
5
fou------>
*/
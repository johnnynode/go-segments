package main

import "fmt"

func main() {
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	fmt.Println();

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	fmt.Println();

	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println();

	for k := range kvs {
		fmt.Println("key", k)
	}
	fmt.Println();

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

/*
$ go run src/range/range.go
sum: 9

index: 1

a -> apple
b -> banana

key b
key a

0 103
1 111
*/
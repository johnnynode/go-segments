package main

import "fmt"

func main() {
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) // sum: 9

	for i, num := range nums {
		if num == 3{
			fmt.Println("index:", i) // index: 1
		}
	}

	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // a -> apple ↵ b -> banana
	}

	for k := range kvs {
		fmt.Println("key:",k) // key: a ↵ key: b
	}

	for i, c := range "go" {
		fmt.Println(i, c) // 0 103 ↵ 1 111
	}
}
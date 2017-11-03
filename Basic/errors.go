package main

import (
	"errors"
	"fmt"
)

// 参数为42 返回异常
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	// 否则返回原值+3
	return arg + 3, nil
}

// 自定义一个arg参数 struct
type argError struct {
	arg  int
	prob string
}

// 给argError 定义一个Error方法 返回类型为string
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant' work with it"}
	}
	return arg + 3, nil
}

func main() {
	// for range 循环一个数组
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	// 强制转换error成argError类型 注意这里的ok
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

/*

$ go run Basic/errors.go
f1 worked: 10
f1 failed: Can't work with 42
f2 worked: 10
f2 failed: 42 - cant' work with it
42
cant' work with it

*/
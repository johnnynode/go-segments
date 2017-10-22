package main

import "fmt"

func main() {
    i := 1
    for i <=3 {
        fmt.Println(i) // 换行输出 1 2 3
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j) // 换行输出 7 8 9
    }

    // 无限循环
    /*
    for {
        fmt.Println("loop");
    }
    */

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n) // 换行输出1 3 5
    }
}
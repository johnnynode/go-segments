package main

import "fmt"
import "math"

const s string = "constant"

func main() {
    fmt.Println(s) // constant
    const n = 500000000
    const d = 3e20 / n
    fmt.Println(d) // 6e+11
    fmt.Println(int64(d)) // 600000000000
    fmt.Println(math.Sin(n)) // -0.28470407323754404
}
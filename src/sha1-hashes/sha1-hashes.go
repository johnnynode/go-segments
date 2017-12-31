package main

import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1  this string"
	h := sha1.New()

	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

/*
$ go run src/sha1-hashes/sha1-hashes.go
sha1  this string
83e2a64a24675036d27ff899753280baa70aed8e
*/

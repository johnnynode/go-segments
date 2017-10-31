package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./Basic/hello-world.go")
	check(err)
	fmt.Printf(string(dat))

	f, err := os.Open("./Basic/hello-world.go")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)

	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	f.Close()

}

/*

$ go run Basic/reading-files.go
package main

import "fmt"

func main() {
        fmt.Println("hello world")
}
5 bytes: packa
2 bytes @ 6: e
5 bytes: packa

*/
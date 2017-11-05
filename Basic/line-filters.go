package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}

/*

$ echo 'hello' > ./Basic/line-filters-test.txt
$ echo 'filter' >> ./Basic/line-filters-test.txt
$ cat ./Basic/line-filters-test.txt | go run Basic/line-filters.go
HELLO
FILTER

*/

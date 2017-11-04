package main

import (
	s "strings"
	"fmt"
)

var p = fmt.Println

func main() {
	p("Containes: ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "e"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Join([]string{"a", "b"}, "-"))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:    ", s.Replace("foo", "o", "0", 1))
	p("Split:      ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:    ", s.ToLower("TEST"))
	p("ToUpper:    ", s.ToUpper("test"))
	p()

	p("Len:        ", len("hello"))
	p("Char:       ", "hello"[1])
}

/*

$ go run Basic/string-functions.go
Containes:  true
Count:      2
HasPrefix:  true
HasSuffix:  false
Index:      1
Join:       a-b
Repeat:     a-b
Replace:    f00
Replace:     f0o
Split:       [a b c d e]
ToLower:     test
ToUpper:     TEST

Len:         5
Char:        101

*/

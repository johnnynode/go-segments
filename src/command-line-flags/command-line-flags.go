package main

import "flag"
import "fmt"

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

/*
$ go build -o test/command-line-flags src/command-line-flags/command-line-flags.go

$ test/command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

$ test/command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

$ test/command-line-flags -word=opt a1 a2 a3
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3]

$ test/command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

$ test/command-line-flags -h
Usage of test/command-line-flags:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
				a string (default "foo")

$ test/command-line-flags -wat
flag provided but not defined: -wat
Usage of test/command-line-flags:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")

*/

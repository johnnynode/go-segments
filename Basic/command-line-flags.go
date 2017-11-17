package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("foor", false, "a bool")

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
$ go build Basic/command-line-flags.go

$ ./command-line-flags -word=opt =numb=7 -fork -svar=flag
word: optnumb: 42
fork: falsesvar: bar
tail: [=numb=7 -fork -svar=flag]

$ ./command-line-flags -word=opt
word: opt
numb: 42fork: false
svar: bartail: []

$ ./command-line-flags -word=opt a1 a2 a3
word: opt
numb: 42fork: false
svar: bartail: [a1 a2 a3]

$ ./command-line-flags a1 a2 a3 -numb=7
word: foo
numb: 42fork: false
svar: bartail: [a1 a2 a3 -numb=7]

$ ./command-line-flags -h
Usage of ./command-line-flags:
  -foor        a bool
  -numb int        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
				a string (default "foo")
				
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:  -foor
        a bool  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")				
*/
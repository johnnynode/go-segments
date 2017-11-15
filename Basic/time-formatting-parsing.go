package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	
	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)

	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}

/*
$ go run Basic/time-formatting-parsing.go
2017-11-15T08:36:33+08:00
2012-11-01 22:08:41 +0000 +0000
8:36AM
Wed Nov 15 08:36:33 2017
2017-11-15T08:36:33.703706+08:00
0000-01-01 20:41:00 +0000 UTC
2017-11-15T08:36:33-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
*/
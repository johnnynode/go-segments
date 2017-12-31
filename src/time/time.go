package main

import "fmt"
import "time"

func main() {
	p := fmt.Println
	
	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		p(then)
		
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

/*
$ go run src/time/time.go
2017-12-31 19:38:17.846373 +0800 CST m=+0.000348623
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
71175h3m19.194985763s
71175.05533194049
4.270503319916429e+06
2.5623019919498578e+08
256230199194985763
2017-12-31 11:38:17.846373 +0000 UTC
2001-10-05 05:31:39.456401474 +0000 UTC
*/
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme) // postgres
	fmt.Println(u.User) // user:pass
	fmt.Println(u.User.Username()) // user
	p, _ := u.User.Password()
	fmt.Println(p) // pass
	fmt.Println(u.Host) // host.com:5432
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host) // host.com
	fmt.Println(port) // 5432
	fmt.Println(u.Path) // /path
	fmt.Println(u.Fragment) // f
	fmt.Println(u.RawQuery) // k=v
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m) // map[k:[v]]
	fmt.Println(m["k"][0]) // v
}

/*

$ go run basic/url-parsing.go
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v

*/
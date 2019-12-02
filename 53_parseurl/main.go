package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)          // postgres://user:pass@host.com:5432/path?k=v#f
	fmt.Println(u.String()) // postgres://user:pass@host.com:5432/path?k=v#f
	fmt.Println(u.Scheme)   // postgres
	fmt.Println(u.Path)     // /path
	fmt.Println(u.Fragment) // f
	fmt.Println(u.RawQuery) // k=v
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)         // map[k:[v]]
	fmt.Println(m["k"][0]) // v

	fmt.Println(u.Host)       // host.com:5432
	fmt.Println(u.Hostname()) // host.com
	fmt.Println(u.Port())     // 5432

	fmt.Println(u.User)            // user:pass
	fmt.Println(u.User.String())   // user:pass
	fmt.Println(u.User.Username()) // user
	fmt.Println(u.User.Password()) // pass true

}

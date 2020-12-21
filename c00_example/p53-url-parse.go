package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main(){
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u,err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) // postgres
	fmt.Println(u.User) // user:pass
	fmt.Println(u.User.Username()) // user

	password, _ := u.User.Password()
	fmt.Println(password) // pass

	fmt.Println(u.Host) // host.com:5432
	h := strings.Split(u.Host,":")
	fmt.Println(h[0]) // host.com
	fmt.Println(h[1]) // 5432

	fmt.Println(u.Path) // /path
	fmt.Println(u.Fragment) // f

	fmt.Println(u.RawQuery) // k=v
	query, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(query) // map[k:[v]]
	fmt.Println(query["k"][0]) // v

}

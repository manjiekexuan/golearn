package main

import (
	"fmt"
	"golearn/ch6/url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))

	m = nil

	fmt.Println(m.Get("item"))
	m.Add("item", "3")

}

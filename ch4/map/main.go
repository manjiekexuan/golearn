package main

import "fmt"

func main() {
	mapAbc := make(map[string]int)

	mapAbc["abc"] = 1
	mapAbc["bbc"] = 2

	for key, value := range mapAbc {
		fmt.Println(key, value)
	}
}

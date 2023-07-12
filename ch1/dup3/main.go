package main

import (
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		fmt.Println(filename)
	}
	fmt.Println(counts)
}

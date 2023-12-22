package main

import (
	"fmt"
	"os"
)

// 相对于echo1的写法优化，使用range来遍历os.Args，而且使用了[1:]slice结构
func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

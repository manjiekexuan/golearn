package main

import (
	"fmt"
	"os"
	"strings"
)

// 相对于echo1和echo2 的+=的""方式，当数据量比较大消耗会很大，使用strings包中的Join的方法。
func main() {
	fmt.Println(strings.Join(os.Args[1:], ""))
}

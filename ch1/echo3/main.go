package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], ""))
	//“练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字”
	//fmt.Println(strings.Join(os.Args[0:], ""))

	// 练习1.2 “修改echo程序，使其打印每个参数的索引和值，每个一行”
	//for k, v := range os.Args[1:] {
	//	println(k, v)
	//}

}

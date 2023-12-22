package main

import (
	"fmt"
	"os"
)

// 使用os包中的Args来接收命令
// 例子: go run main.go 命令 参数
// 结果: 命令 参数
// 结论也就是说os.Args 是个sliec的结构  os.Args[1] = 命令 os.Args[2] = 参数 # os.Args[0] 是命令本身的名字
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

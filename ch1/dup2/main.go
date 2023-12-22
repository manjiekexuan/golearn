package main

import (
	"bufio"
	"fmt"
	"os"
)

// 能使用调试尽量使用debug来进行调试。
// 发现报错  version of Delve is too old for Go version 1.21.0 (maximum supported version 1.19) 这里使用
// 1.go get -u github.com/go-delve/delve/cmd/dlv
// 2.设置Goland 帮助-自定义编辑 dlv.path = /Users/xueye/go/bin/dlv
func main() {
	counts := make(map[string]int)
	files := []string{"/Users/xueye/code/golearn/ch1/dup2/dup2.txt"}
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

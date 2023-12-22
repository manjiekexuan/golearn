package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio包 bufio 用来帮助处理 I/O 缓存。 我们将通过一些示例来熟悉其为我们提供的：Reader, Writer and Scanner 等一系列功能
func main() {
	// 这里有个很好的说法就是map里面的键值 只要可以用==进行比较就可以。
	counts := make(map[string]int)
	//继续来看 bufio 包，它使处理输入和输出方便又高效。Scanner 类型是该包最有用的特性之一，它读取输入并将其拆成行或单词；通常是处理行形式的输入最简单的方法。
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		// 在linux 和 windows 里使用ctrl+d 就可以终止，打印结果。
		if input.Text() == "bbc" {
			break
		}
	}
	for line, n := range counts {
		if n > 1 {
			//%d          十进制整数
			//%x, %o, %b  十六进制，八进制，二进制整数。
			//%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
			//%t          布尔：true或false
			//%c          字符（rune） (Unicode码点)
			//%s          字符串
			//%q          带双引号的字符串"abc"或带单引号的字符'c'
			//%v          变量的自然形式（natural format）
			//%T          变量的类型
			//%%          字面上的百分号标志（无操作数）

			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// 因为是标准输入OS.Stin 所以需要先编译执行，但是执行报错golang syntax error near unexpected token `newline'
	// 这是因为package 要在main下执行，我写的dup1 。说明可执行文件 要在main下才能执行 ？
}

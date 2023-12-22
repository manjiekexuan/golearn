package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	urls := []string{"www.google.com"}
	for _, url := range urls {
		//例2 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		//例3 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码
		fmt.Printf(resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		//原
		//b, err := io.ReadAll(resp.Body)

		//例1
		out, err := os.Create("/Users/xueye/code/golearn/ch1/dup3/dup3.txt")
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Printf("copy 报错")
		}
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", url, err)
			os.Exit(1)
		}

	}

}

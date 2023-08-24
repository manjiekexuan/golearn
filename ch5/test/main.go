package main

import (
	"fmt"
	"os"
)

//func main() {
//	var rmdirs []func()
//	for _, dir := range tempDirs() {
//		dir := dir
//		fmt.Println("处理", dir)
//		os.MkdirAll(dir, 0755)
//		rmdirs = append(rmdirs, func() {
//			os.RemoveAll(dir)
//		})
//	}
//}

func tempDirs() []string {

	returnSlice := []string{
		"/Users/g01d-01-0221/go/golearn/ch5/test/1", "/Users/g01d-01-0221/go/golearn/ch5/test/2", "/Users/g01d-01-0221/go/golearn/ch5/test/3",
	}
	return returnSlice
}

func main() {
	var rmdirs []func()
	dirs := tempDirs()
	for i := 0; i < len(dirs); i++ {
		os.MkdirAll(dirs[i], 0755)
		rmdirs = append(rmdirs, func() {
			fmt.Println(dirs[i])
			os.RemoveAll(dirs[i])
		})
	}
}

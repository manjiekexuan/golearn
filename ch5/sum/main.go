package main

import (
	"fmt"
	"os"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

	linenum, name := 12, "count"
	errorf(linenum, "undefined:%s", name)
}

func f(...int) {}

func g([]int) {}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d:", linenum)
	fmt.Fprintf(os.Stderr, format, args)
	fmt.Fprintln(os.Stderr)
}

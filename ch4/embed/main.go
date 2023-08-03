package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	//w := Wheel{Circle{Point{8, 8}, 5}, 20}
	w := Wheel{Circle: Circle{Center: Point{X: 8, Y: 8}, Radius: 5}, Spokes: 20}
	w.Circle.Center.X = 48
	fmt.Printf("%#v\n", w)

}

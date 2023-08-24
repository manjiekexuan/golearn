package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{X: 20, Y: 30}
	q := Point{X: 40, Y: 60}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	perim := Path{
		{1, 1}, {5, 1}, {5, 4}, {1, 1},
	}
	fmt.Println(perim.Distance())
}

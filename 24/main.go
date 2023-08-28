package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPointer(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p.x-p2.x, 2) + math.Pow(p.y-p2.y, 2))
}

func main() {

	p1 := NewPointer(2.4, -3.5)
	p2 := NewPointer(5, 3.1)

	fmt.Println(p1.distance(p2))
}

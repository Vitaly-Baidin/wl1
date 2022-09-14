package main

import (
	"fmt"
	"math"
)

// TODO| Разработать программу нахождения расстояния между двумя точками,
// TODO| которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) CalcDistance(target *Point) float64 {
	x := target.x - p.x
	y := target.y - p.y

	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func main() {
	p1 := NewPoint(12, 12)
	p2 := NewPoint(14, 14)

	fmt.Println(p1.CalcDistance(p2))
}

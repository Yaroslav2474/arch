package main

import (
	"fmt"
	"math"
)

type rect struct {
	height int
	width  int
}

func (r rect) perim() int {
	return (r.height + r.width) * 2
}

func (r rect) area() int {
	return r.height * r.width
}

type shape struct {
	radius int
}

type circle interface {
	perim() float64
	area() float64
}

func (s shape) perim() float64 {
	return (2 * math.Pi * float64(s.radius))
}

func (s shape) area() float64 {
	return (math.Pi * math.Pow(float64(s.radius), 2))
}

func printZ(circle circle) {
	fmt.Printf("Площадь: %.1f\nПериметр: %.1f\n", circle.area(), circle.perim())
}

func main() {
	cir := shape{
		radius: 10,
	}
	printZ(cir)

	kvadrat := rect{
		height: 20,
		width:  20,
	}

	kvadrat := rect{}
	kvadrat.height = 20
	kvadrat.width = 20

	fmt.Printf("Площадь: %d\nПериметр: %d\n", kvadrat.area(), kvadrat.perim())
}

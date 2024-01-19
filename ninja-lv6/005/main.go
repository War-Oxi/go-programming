package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}

func main() {
	myCircle := circle{
		radius: 12.345,
	}

	mySquare := square{
		length: 15,
	}
	info(myCircle)
	info(mySquare)
}

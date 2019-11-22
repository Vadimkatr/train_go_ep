package figure

import (
	"math"
)

type Figure interface {
	Area() float64
	Perimeter() float64
}

type Point struct {
	X, Y int
}

type Square struct {
	Start Point
	A, B  float64
}

func (s Square) Area() float64 {
	return s.A * s.B
}

func (s Square) Perimeter() float64 {
	if s.A == 0 || s.B == 0 {
		return 0
	}
	return (s.A + s.B) * 2
}

type Circle struct {
	Center Point
	R      float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

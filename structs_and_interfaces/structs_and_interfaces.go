package structs_and_interfaces

import (
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// circumference btw
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func (t Triangle) Perimeter() float64 {
	return t.Base + (t.Height * 2)
}

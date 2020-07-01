package interfaces

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle sruct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Triangle sruct
type Triangle struct {
	Base   float64
	Height float64
}

// Area return float64
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area return float64
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area return float64
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

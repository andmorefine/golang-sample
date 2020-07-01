package structs

import "math"

// Rectangle sruct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area return float64
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area return float64
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

package main

import "math"

// Rectangle struct to capture the proper dimensions.
type Rectangle struct {
	width  float64
	height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Shape is the interface of the Area function
type Shape interface {
	Area() float64
}

// Returns perimeter of a rectangle
func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

// Returns area of a rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

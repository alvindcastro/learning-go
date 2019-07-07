package structs

import "math"

//Shape - Interface
type Shape interface {
	Area() float64
}

//Rectangle - Struct
type Rectangle struct {
	Width  float64
	Height float64
}

//Area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Circle - Struct
type Circle struct {
	Radius float64
}

//Area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Triangle - Struct
type Triangle struct {
	Base   float64
	Height float64
}

//Area of Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

//Perimeter - Initial code
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area - Initial code
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

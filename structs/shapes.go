package structs

//Rectangle - struct
type Rectangle struct {
	Width  float64
	Height float64
}

//Perimeter - Initial code
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area - Initial code
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

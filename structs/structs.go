package structs

const PI = 3.1416

type Rectangle struct {
	Length float64
	Width  float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Perimeter() float64 {
	return 2 * PI * c.Radius
}

func (c Circle) Area() float64 {
	return PI * (c.Radius * c.Radius)
}

func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return 0.0
}

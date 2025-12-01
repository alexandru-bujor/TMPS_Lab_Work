package Lab_0

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }

type Square struct{ Side float64 }

func (s Square) Area() float64 { return s.Side * s.Side }

func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

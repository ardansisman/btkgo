package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectange struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectange) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func calculate(s shape) {
	fmt.Println("Şeklin alanı:", s.area())

}
func Demo1() {
	r := rectange{width: 10, height: 6}
	calculate(r)
	c := circle{radius: 10}
	calculate(c)
}

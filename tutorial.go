package main

import (
	"fmt"
)

type shape interface {
	area() float64
}

type trapezium struct {
	basea  float64
	baseb  float64
	height float64
}
type rhombus struct {
	diagonala float64
	diagonalb float64
}

func (t trapezium) area() float64 {
	return ((t.basea + t.baseb) / 2) * t.height
}
func (r rhombus) area() float64 {
	return ((r.diagonala * r.diagonalb) / 2)
}

func Result(s shape) float64 {
	return s.area()
}

func main() {
	r1 := rhombus{6, 4}
	t1 := trapezium{8, 6, 4}
	shapes := []shape{r1, t1}

	for _, shape := range shapes {
		fmt.Println(Result(shape))
	}
}

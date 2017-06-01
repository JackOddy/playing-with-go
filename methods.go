package main

import "math"

type Circle struct {
	r, x, y float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	w, l float64
}

func (r *Rectangle) area() float64 {
	return r.l * r.w
}
func main() {
	c := Circle{5, 0, 0}
	ca := c.area()
	println(ca)
	r := Rectangle{5, 10}
	ra := r.area()
	println(ra)
}

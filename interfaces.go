package main

import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) (area float64) {
	for _, s := range shapes {
		area += s.area()
	}
	return
}

func totalPerimeter(shapes ...Shape) (perimeter float64) {
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return
}

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	w, l float64
}

func (r *Rectangle) area() float64 {
	return r.w * r.l
}

func (r *Rectangle) perimeter() float64 {
	return (2 * r.w) + (2 * r.l)
}

type Triangle struct {
	s float64
}

func (t *Triangle) area() float64 {
	return (t.s * t.s) / 2
}

func (t *Triangle) perimeter() float64 {
	return t.s * 3
}

func main() {
	c := Circle{10}
	r := Rectangle{5, 10}
	t := Triangle{10}
	println(totalArea(&c, &r, &t))
	println(totalPerimeter(&c, &r, &t))

}

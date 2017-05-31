package main

func main() {
	x := 1
	y := 2
	swap(&x)(&y)
	println(x, y)
}

func swap(x *int) func(*int) {
	xn := *x
	return func(y *int) {
		yn := *y
		*x = yn
		*y = xn
	}

}

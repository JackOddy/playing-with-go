package main

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	println(x, y)
}

func swap(x, y *int) {
	defer func(n int) { *x = n }(*y)
	*y = *x
}

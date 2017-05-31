package main

import "fmt"

func main() {
	n := curry()(1)
	fmt.Println(n)
}

func curry() func(int) int {
	return sauce
}

func sauce(arg int) int {
	return arg * 2
}

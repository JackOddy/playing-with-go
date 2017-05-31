package main

import "fmt"

func main() {
	add := func(y int) func(int) int {
		return func(x int) int {
			return x + y
		}
	}
	fmt.Println(add(1)(1))
}

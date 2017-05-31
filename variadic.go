package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(add(xs...))
}

func add(args ...int) (total int) {
	for _, n := range args {
		total += n
	}
	return
}

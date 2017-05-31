package main

import "fmt"

func main() {
	x := [5]float64{
		98,
		99,
		87,
		76,
		73,
	}
	var total float64
	length := len(x)
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(length))
}

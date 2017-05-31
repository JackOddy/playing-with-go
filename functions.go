package main

import "fmt"

func main() {

	arr := []float64{
		1,
		1,
		1,
		1,
		1,
	}
	avg, total := average(arr)
	fmt.Println(avg)
	fmt.Println(total)
}

func average(xs []float64) (avg, total float64) {
	for _, n := range xs {
		total += n
	}
	avg = total / float64(len(xs))
	return
}

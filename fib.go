package main

import "fmt"

func main() {
	x := getfib(20)
	fmt.Println(x)
}

func getfib(x int) []int {
	slice := []int{0, 1, 1}
	return fib(slice, x-3)
}

func fib(sl []int, n int) []int {
	length := len(sl)
	if n == 0 {
		return sl
	}
	nextNum := sl[length-1] + sl[length-2]
	return fib(append(sl, nextNum), n-1)
}

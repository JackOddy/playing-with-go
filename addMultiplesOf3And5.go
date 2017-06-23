package main

func main() {
	println(Multiple3And5(10))
}
func Multiple3And5(number int) int {
	return findMultiples(number-1, 0)
}

func findMultiples(number, acc int) int {
	if number == 0 {
		return acc
	}

	if isMultiple(number) {
		acc += number
	}

	return findMultiples(number-1, acc)
}

func isMultiple(number int) (res bool) {
	if number%3 == 0 {
		res = true
	}
	if number%5 == 0 {
		res = true
	}
	return
}

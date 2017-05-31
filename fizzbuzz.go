package main

import (
	"fmt"
	"time"
)

func main() {
	tb := makeTimestamp()
	for i := 1; i <= 100; i++ {
		x := ""
		if i%3 == 0 {
			x += "fizz"
		}
		if i%5 == 0 {
			x += "buzz"
		}
		if x == "" {
			fmt.Println(i)
		} else {

			fmt.Println(x)
		}
	}
	ta := makeTimestamp()
	time := ta - tb
	fmt.Println(time)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Nanosecond)
}

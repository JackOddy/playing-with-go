package main

import (
	"fmt"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
	}
}

func printer(c chan string) {

	fmt.Println(<-c)
}

func main() {
	var c chan string = make(chan string, 20)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

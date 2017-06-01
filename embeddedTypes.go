package main

type Animal struct {
	name string
	Dog
}

type Dog struct {
	noise string
}

func (d *Dog) talk() {
	println(d.noise)
}

func main() {
	d := Animal{"Jack", Dog{"Bork"}}
	d.talk()
}

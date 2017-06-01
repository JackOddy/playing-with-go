package main

type Person struct {
	name, surname string
	age           int
}

func main() {
	p := Person{"Jack", "Oddy", 26}
	printPerson(p)
}

func printPerson(p Person) {
	println(p.name)
	println(p.surname)
	println(p.age)
}

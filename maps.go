package main

import "fmt"

func main() {
	x := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
	}
	if name, ok := x["H"]["name"]; ok {
		fmt.Println(name, ok)
	}
}

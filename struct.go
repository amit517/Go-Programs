package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	var p1 Person

	p1.name = "Amit"
	p1.age = 26

	fmt.Println(p1.name)
}

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

	p2 := new(Person)

	/* 	p2.name = "Humane"
	   	p2.age = 0 */
	fmt.Println(p2)

	//p3 :=Person(name:"joe",age:26)
}

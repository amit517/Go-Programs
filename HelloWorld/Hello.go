package main

import "fmt"

type Grades int

const (
	A Grades = iota
	B
	C
	D
	F
)

func main() {
	var x = 2

	if x < 5 {
		fmt.Println("x smaller")
	} else {
		fmt.Println("x grater")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("x Hello")
	}

	switch x {
	case 1:
		{

		}
	case 2:
		{

		}
	}

}

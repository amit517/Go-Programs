package main

import (
	"fmt"
)

func main() {
	/* var x [5]int
	x[0] = 2
	fmt.Printf("%d", x[0]) */

	x := [...]int{1, 2, 3, 4}

	/* for i := 0; i < len(x); i++ {
		fmt.Println("%d", x[i])
	} */

	for i, v := range x {
		fmt.Println("ind %d, val %d", i, v)
	}
}

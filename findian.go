package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string
	fmt.Printf("Please enter a string : ")
	fmt.Scanln(&inputString)
	var lowerCase = strings.ToLower(inputString)
	var isFirstCharContainI = false
	var isStringContainA = false
	var isLastCharContainN = false

	for pos, char := range lowerCase {
		if pos == 0 && char == 'i' {
			isFirstCharContainI = true
		}
		if char == 'a' {
			isStringContainA = true
		}
		if pos == (len([]rune(lowerCase))-1) && char == 'n' {
			isLastCharContainN = true
		}
	}

	if isFirstCharContainI && isStringContainA && isLastCharContainN {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}

package main

import (
	"fmt"
)

func main() {
	/* 	var idMap map[string]int
	   	idMap = make(map[string]int) */

	idMap := map[string]int{
		"joe": 420}

	fmt.Println(idMap["joe"])

	idMap["amit"] = 265
	fmt.Println("my id is", idMap["amit"])
	/* 	delete(idMap, "amit")
	   	fmt.Println("my id is", idMap["amit"]) */

	id, p := idMap["joe"]
	fmt.Println("my id is", id)
	fmt.Println("the value of p is", p)

	for key, value := range idMap {
		fmt.Println("key ", key, "value ", value)
	}
}

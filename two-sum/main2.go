package main

import "fmt"

func MyFunction() {
	fmt.Println("here")
	myMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	if value, exists := myMap["banana"]; exists {
		fmt.Println("Value found:", value)
	} else {
		fmt.Println("Key not found")
	}
}

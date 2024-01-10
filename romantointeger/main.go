package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "III"

	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	number := 0
	rNumerals := make(map[string]int)
	rNumerals["I"] = 1
	rNumerals["V"] = 5
	rNumerals["X"] = 10
	rNumerals["L"] = 50
	rNumerals["C"] = 100
	rNumerals["D"] = 500
	rNumerals["M"] = 1000
	sAsCharacters := strings.Split(s, "")
	for i := 0; i < len(sAsCharacters); i++ {
		if val, exists := rNumerals[sAsCharacters[i]]; exists {
			if i < len(sAsCharacters)-1 {
				nextVal := rNumerals[sAsCharacters[i+1]]
				if val < nextVal {
					number += nextVal - val
					if i < len(sAsCharacters) {
						i++
					}
					continue
				} else {
					number += val
					continue
				}
			} else {
				number += val
			}
		}
	}

	return number
}

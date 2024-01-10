package main

import (
	"fmt"
)

var nums = []int{2, 7, 11, 15, 3, 2, 4}

func main() {
	fmt.Println(twoSum(nums, 6))
	MyFunction()
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, v := range nums {
		remainder := target - v

		if j, found := seen[remainder]; found {
			return []int{j, i}
		}

		seen[v] = i
	}

	return nil // No solution found
}

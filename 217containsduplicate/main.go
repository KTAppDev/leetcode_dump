package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 9}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

package main

import (
	"fmt"
	"strings"
)

var strs = []string{"flower", "flight", "flow"}

func main() {
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	type shortestStr struct {
		str    string
		length int
		index  int
	}
	sS := shortestStr{"", 200, 0}
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < sS.length {
			sS = shortestStr{strs[i], len(strs[i]), i}
		}
	}
	count := 0
	for {
		for j, k := range strs {
			if sS.index == j && sS.index != len(strs)-1 {
				continue
			}
			fmt.Println(k, sS.str[0:sS.length-count])
			if !strings.HasPrefix(k, sS.str[0:sS.length-count]) {
				count++
				break
			} else {
				if j != len(strs)-1 {
					continue
				} else {
					return sS.str[0 : sS.length-count]
				}
			}

		}
	}
}

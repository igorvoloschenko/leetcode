package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Slice(strs, func(i int, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	for i := 1; i < len(strs[0])+1; i++ {
		for _, str := range strs {
			if strs[0][:i] != str[:i] {
				return strs[0][:i-1]
			}
		}
	}
	return strs[0]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Printf("longestCommonPrefix(%v) -> %s\n", strs, longestCommonPrefix(strs))
}

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	store := make(map[rune]int)
	startIndex := 0
	for i, ch := range s {
		if v, ok := store[ch]; ok && v >= startIndex {
			startIndex = v + 1
		}
		store[ch] = i
		if maxLen < i-startIndex+1 {
			maxLen = i - startIndex + 1
		}
	}
	return maxLen
}

func main() {
	tests := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"dvdf",
		"tmmzuxt",
	}

	for _, s := range tests {
		fmt.Println(lengthOfLongestSubstring(s))
	}
}

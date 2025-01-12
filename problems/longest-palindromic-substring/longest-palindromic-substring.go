package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	runes := []rune(s)
	if len(runes) < 2 {
		return s
	}

	var (
		begin, end, nextPosition int
	)

	result := []rune{runes[0]}
	palindromeState := 0

	for i := 1; i < len(runes); i++ {
		if palindromeState < 2 && runes[i] == runes[i-1] {
			if palindromeState == 0 {
				begin = i - 1
				nextPosition = i
				palindromeState = 1
			}
			end = i + 1
			continue
		}

		if i > 1 && palindromeState == 0 && runes[i] == runes[i-2] {
			begin = i - 2
			end = i + 1
			nextPosition = i
			palindromeState = 2
			continue
		}

		if palindromeState > 0 && begin > 0 && runes[i] == runes[begin-1] {
			begin--
			end = i + 1
			palindromeState = 2
			continue
		}

		if palindromeState > 0 {
			if len(result) < len(runes[begin:end]) {
				result = runes[begin:end]
			}
			palindromeState = 0
			i = nextPosition
		}
	}

	if palindromeState > 0 {
		if len(result) < len(runes[begin:end]) {
			result = runes[begin:end]
		}
	}

	return string(result)
}

func main() {
	tests := []string{
		"",
		"a",
		"asdasdsasd",
		"babad",
		"abababa",
		"cbbd",
		"aacabdkacaa",
		"abbcccbbbcaaccbababcbcabca",
		"bananas",
	}

	for _, test := range tests {
		fmt.Printf(
			"longestPalindrome('%s') -> '%s'\n",
			test,
			longestPalindrome(test),
		)
	}
}

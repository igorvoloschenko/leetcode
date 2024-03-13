package main

import "fmt"

func longestPalindrome(s string) string {
	return ""
}

func main() {
	tests := []string{
		"babad",
		"cbbd",
	}

	for _, test := range tests {
		fmt.Printf(
			"longestPalindrome('%s') -> '%s'",
			test,
			longestPalindrome(test),
		)
	}
}

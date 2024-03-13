package main

import "fmt"

func romanToInt(s string) int {
	romanInt := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	out := 0
	last := romanInt['M']
	for _, v := range s {
		if romanInt[v] <= last {
			out += romanInt[v]
			last = romanInt[v]
		} else {
			out += romanInt[v] - (last)*2
		}
	}
	return out
}

func main() {
	s := "MCMXCIV"
	fmt.Printf("romanToInt('%s') -> %d\n", s, romanToInt(s))
}

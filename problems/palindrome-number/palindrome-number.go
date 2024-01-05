package main

import "fmt"

func reverse(x int) int {
	res := 0
	for x != 0 {
		res *= 10
		res += x % 10
		x /= 10
	}

	return res
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == reverse(x)
}

func main() {
	test_nums := []int{
		11,
		121,
		122,
		344,
		555,
		6666,
		1,
		15,
		-121,
	}

	for _, num := range test_nums {
		fmt.Printf("isPalindrome(%d) = %v\n", num, isPalindrome(num))
	}

}

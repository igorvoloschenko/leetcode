package main

import "fmt"

func twoSum(nums []int, target int) []int {
	store := make(map[int]int)
	for i, num := range nums {
		if v, ok := store[target-num]; !ok {
			store[num] = i
		} else {
			return []int{v, i}
		}
	}
	return []int{}
}

func main() {
	nums, target := []int{3, 2, 4}, 6

	fmt.Printf("twoSum(%v, %d) -> %v\n", nums, target, twoSum(nums, target))
}

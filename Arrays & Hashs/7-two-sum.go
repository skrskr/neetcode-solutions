package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	///// time complexity: O(n^2) 
	//// Brute force
	// for i := 0; i < len(nums); i++ {
	// 	for j := i+1; j < len(nums); j++ {
	// 		if nums[i] + nums[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// return []int{}

	// // Better solution
	// // time complexity: O(n)
	// // memory complexity: O(n)
	diffMap := make(map[int]int)
	for pos, num := range nums {

		if val, ok := diffMap[target - num]; ok {
			return []int{val, pos}
		}
		diffMap[num] = pos
	}
	return []int{}
}

func main() {
	nums := []int{2,5,5,7}
	target := 9
	fmt.Println(twoSum(nums, target))
}
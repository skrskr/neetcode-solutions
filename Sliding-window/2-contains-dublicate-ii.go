package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	//// Two pointers but failed on case [99,99], k = 2
	// l, r := 0, k-1
	// for ; r < len(nums); {
	// 	if nums[l] == nums[r] {
	// 		return true
	// 	}
	// 	l++
	// 	r++
	// }
	// return false
	window := map[int]bool{}
	l := 0

	for r := range(nums) {
		if r - l > k {
			delete(window, nums[l])
			l++
		}
		if window[nums[r]] == true {
			return true
		}
		window[nums[r]] = true
	}

	return false
}

func main() {
	k:= 2
	nums := []int{1,2,3,1,2,3}
	fmt.Println(containsNearbyDuplicate(nums, k))
}
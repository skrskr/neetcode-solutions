package main

import (
	"fmt"
)
func findMin(nums []int) int {
    l, r := 0, len(nums) - 1
	// Mathatically is the same
	// mid := (l + r) / 2
	//// is preferred to avoid overflow.
	// mid := l + (r - l) / 2  
	
	for l < r {
		mid := l + (r - l) / 2
		
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

func main() {
	nums := []int{11,13,15,17}
	fmt.Println(findMin(nums))
}
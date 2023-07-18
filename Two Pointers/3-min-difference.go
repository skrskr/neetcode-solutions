package main

import (
	"fmt"
	"math"
	"sort"
)


func minimumDifference(nums []int, k int) int {
	if len(nums) == 1{
		return 0
	}

	// sort array
	sort.Ints(nums)
	min:= math.MaxInt
	for l, r := 0, k - 1;r < len(nums); l, r = l + 1, r+1 {
		if min > (nums[r] - nums[l]) {
			min = nums[r] - nums[l]
		}
	}
	return min
}

func main() {
	k:=2
	nums := []int{7,4,9,1}
	fmt.Println(minimumDifference(nums, k))
}
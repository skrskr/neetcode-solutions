package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	l, total, min := 0,0, 100000000
	
	for r := range(nums) {
		total += nums[r]

		for total >= target {
			total -= nums[l]
			if r - l + 1 < min {
				min = r - l + 1
			}
			l++
		}
	}
	if min == 100000000 {
		return 0
	}
	return min
}

func main() {
	target := 7
	nums := []int{2,3,1,2,4,3}
	fmt.Println(minSubArrayLen(target, nums))
}
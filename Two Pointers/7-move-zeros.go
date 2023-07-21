package main

import (
	"fmt"
)

func moveZeroes(nums []int)  {
    l, r := 0, 1
	for ; r < len(nums); {
		if nums[r] != 0 && nums[l] == 0 {
			nums[l], nums[r] = nums[r], nums[l]
		}
		if nums[l] != 0 {
			l++
		}
		r++
	}
}

func main() {
	nums := []int{1,2}
	moveZeroes(nums)
	fmt.Println(nums)
}
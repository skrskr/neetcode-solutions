package main

import (
	"fmt"
)
func search(nums []int, target int) int {
    low, high := 0, len(nums) - 1

    for low <= high {
        mid := (low + high) / 2

        if nums[mid] == target {
            return mid
        }

        if nums[low] <= nums[mid] {
            if nums[low] <= target && target < nums[mid] {
                high = mid - 1
            } else {
                low = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[high] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }
    }

    return -1
}

func main() {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	fmt.Println(search(nums, target))
}
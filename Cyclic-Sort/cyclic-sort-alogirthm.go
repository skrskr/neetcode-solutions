package main

import "fmt"



func cyclicSort(nums []int) []int {
	i := 0

	for i < len(nums) {
		indexOfNum := nums[i] - 1
		if nums[i] != nums[indexOfNum] {
			nums[i], nums[indexOfNum] = nums[indexOfNum], nums[i]
		} else {
			i++
		}
	}

	return nums
}


func main() {
	nums := []int{4,2,1,3,5}
	fmt.Println(cyclicSort(nums));
}
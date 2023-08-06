package main

import "fmt"

func findDuplicates(nums []int) []int {
    i := 0
	for i < len(nums) {
		j := nums[i] - 1

		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}

	dublicateArray := []int{}
	for i:= 0; i < len(nums); i++ {
		if nums[i] != i +1 {
			dublicateArray = append(dublicateArray, nums[i])
		}
	}
	return dublicateArray
}

func main() {
	nums := []int{1,1,2}
	fmt.Println(findDuplicates(nums));
}
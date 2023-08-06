package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
    i := 0

	for i < len(nums) {
		j := nums[i] - 1

		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}

	missingNums := []int{}
	for i:= 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			missingNums = append(missingNums, i+1)
		}
	}
	return missingNums
}

func main() {
	nums := []int{1, 1}
	fmt.Println(findDisappearedNumbers(nums));
}
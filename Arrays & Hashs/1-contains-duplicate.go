package main

import "fmt"

func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)

	for _, num := range nums {
		_, ok := mp[num]
		if ok {
			return true
		}
		mp[num] = true
	}
	return false
}


func main() {
	nums := []int{1, 1}
	res := containsDuplicate(nums)
	fmt.Println(res)
}
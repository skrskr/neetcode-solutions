package main

import (
	"fmt"
)
func peakIndexInMountainArray(arr []int) int {
	arrLen := len(arr)
	l, r := 0, arrLen -1

	for l <= r {
		mid := l + (r - l) / 2

		if (mid == 0 || arr[mid] > arr[mid -1]) && (mid == arrLen - 1 || arr[mid] > arr[mid + 1]) {
			return mid
		} else if (arr[mid] > arr[mid + 1]) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{10,20,30,40}
	// target := 0
	fmt.Println(peakIndexInMountainArray(nums))
}
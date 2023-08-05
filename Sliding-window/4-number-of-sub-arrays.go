package main

import (
	"fmt"
)

func numOfSubarrays(arr []int, k int, threshold int) int {
    sum, cnt := 0, 0
	for i:= 0; i < k; i++ {
		sum += arr[i]
	}
	avg := sum / k
	if avg >= threshold {
		cnt++
	}
	for i:= k; i < len(arr); i++ {
		sum = sum - arr[i-k] + arr[i]
		avg := sum / k
		if avg >= threshold {
			cnt++
		}	
	}
	return cnt
}

func main() {
	arr := []int{11,13,17,23,29,31,7,5,2,3}
	k, threshold := 3, 5
	fmt.Println(numOfSubarrays(arr, k, threshold))
}
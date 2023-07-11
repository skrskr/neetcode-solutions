package main

import "fmt"

func replaceElements(arr []int) []int {
	//// Brute force solution 1 O(n^2)
    // len := len(arr)
	// for i := 0; i < len; i++ {
		// 	max := -1
	// 	for j := i+1; j < len; j++ {
	// 		if arr[j] > max {
	// 			max = arr[j]
	// 		}
	// 	}
	// 	arr[i] = max
	// }
	// return arr
	
	//// optimal solution 2 O(n)
	len := len(arr)
	max := -1
	for i:= len - 1; i >= 0; i-- {
		newMax := max
		if arr[i] > newMax {
			newMax = arr[i]
		}
		arr[i] = max
		max = newMax
	}
	return arr
}

func main() {
	arr := []int{17,18,5,4,6,1}
	fmt.Println(replaceElements(arr))
}
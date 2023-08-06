package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	///// wrong answer with cyclic sort
	// lastItem := arr[len(arr)-1]
	// for i:= 0; i <= k; i++ {
	// 	arr = append(arr, lastItem)
	// }
	// i := 0
	// for i < len(arr) {
	// 	j := arr[i] - 1
		
	// 	if arr[i] != arr[j] {
	// 		arr[i], arr[j] = arr[j], arr[i]
	// 	} else {
	// 		i++
	// 	}
	// }
	// missingNums := []int{}
	// for i:= 0; i < len(arr); i++ {
	// 	if arr[i] != i +1 {
	// 		missingNums = append(missingNums, i+1)
	// 	}
	// }
	// return missingNums[k-1]
	////////////////////////////////////
	//// binary search
	l, r := 0, len(arr) - 1

	for l <= r {
		mid := (l+r) / 2
		numberOfMissingNumbers := arr[mid] - mid - 1

		if numberOfMissingNumbers < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r + k + 1
}

func main() {
	nums := []int{1,2,3,4}
	k := 2
	fmt.Println(findKthPositive(nums, k));
}
package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
    ///// inital solution
	/*
	* issue with soultion nums1 not reflected num2 if m == 0
	*/
	// if n == 0 {
	// 	return
	// }
	// if m == 0 {
	// 	// nums1 = append(nums1, nums2...)
	// 	// nums1 = nums2
	// 	for i:=0; i<n;i++ {
	// 		nums1 = append(nums1, nums2[i])
	// 	}
	// 	fmt.Println(nums1)
	// 	return
	// }

	// i, j := 0,0

	// for ; i < (n + m) && j < n; {
	// 	if nums1[i] == 0 {
	// 		nums1[i] = nums2[j]
	// 		j++
	// 	} else if nums1[i] > nums2[j] {
	// 		nums1[i], nums2[j] = nums2[j], nums1[i]
	// 	}

	// 	i++
	// }

	last := n + m -1
	m, n = m - 1, n - 1
	for ; m >= 0 && n >= 0; {
		if nums1[m] > nums2[n] {
			nums1[last] = nums1[m]
			m--
		} else {
			nums1[last] = nums2[n]
			n--
		}
		last--
	}
	for ; n >= 0; {
		nums1[last] = nums2[n]
		n--
		last--
	}
}

func main() {
	// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	var nums1 = []int{1,2,3,0,0,0}
	var nums2 = []int{2,5,6}
	m, n := 3,3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)	
}
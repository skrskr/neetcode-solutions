package main

import "fmt"

func getConcatenation(nums []int) []int {
    //// solution 1
	// ans := append(nums, nums...)
    // return ans
    
	//// solution 2
	// ans := make([]int, 0)
    // for i:= 0; i < 2; i++{
    //     for _, n := range nums {
    //         ans = append(ans, n)
    //     }
    // }
    // return ans
	//// solution 3
    numsLen := len(nums)
	ans := make([]int, numsLen, numsLen * 2)
	copy(ans, nums)
	ans = append(ans, nums...)
	return ans

}

func main() {
	
	nums := []int{1, 2, 3, 4}
	fmt.Println(getConcatenation(nums))
}
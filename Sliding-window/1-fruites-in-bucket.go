package main

import (
	"fmt"
)

func totalFruit(fruits []int) int {
	left, res, total := 0, 0, 0
	count := map[int]int{}

	for _, fruit := range(fruits) {
		count[fruit]++
		total++

		for len(count) > 2 {
			f := fruits[left]
			total--
			left++
			count[f] --
			if count[f] == 0 {
				delete(count, f)
			}
		}
		if total > res {
			res = total
		}
	} 
	return res
}

func main() {
	nums := []int{1,2,3,1}
	fmt.Println(totalFruit(nums))
}
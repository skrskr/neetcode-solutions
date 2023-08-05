package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
    l, total, max := 0, 0, 0
	window := map[byte]bool{}

	for r := range(s) {

		for ; window[s[r]] == true; {
			delete(window, s[l])
			l++
		}

		window[s[r]] = true
		
		total = r - l + 1
		if total > max {
			max = total
		}
	}
	return max
}

func main() {

	fmt.Println(lengthOfLongestSubstring("bbb"))
}
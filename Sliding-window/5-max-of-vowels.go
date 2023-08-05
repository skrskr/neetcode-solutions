package main

import (
	"fmt"
)

func maxVowels(s string, k int) int {
    l, total, max := 0, 0, 0
	for r := range(s) {
		if isVowel(s[r]) {
			total++
		}
		if r - l + 1 > k {
			if isVowel(s[l]) {
				total--
			}
			l++
		}

		if total > max {
			max=total
		}
	}

	return max
}

func isVowel(char byte) bool {
	return (char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u')
}

func main() {
	k := 3
	s := "leetcode"
	fmt.Println(maxVowels(s, k))
}
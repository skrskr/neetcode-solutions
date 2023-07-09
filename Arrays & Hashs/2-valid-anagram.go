package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	result := [26]int{0}
	for indx, ch := range s {
		result[ch - 'a']++
		result[t[indx] -'a']--
	}
	for _, val := range result {
		if val != 0 {
			return false
		}
	}
	return true
}

func main() {
	s, t := "anagram", "nagaraf"
	fmt.Println(isAnagram(s, t))
}
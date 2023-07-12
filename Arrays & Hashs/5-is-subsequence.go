package main

import (
	"fmt"
	// "strings"
)

func isSubsequence(s string, t string) bool {
	// //// worng answer with case s ="ab" and t = "baab"
	// sLen := len(s)
	// tLen := len(t)
	// if sLen > tLen {
	// 	return false
	// }
    // newIndex := -1
	// for _, c := range s {
	// 	indx := strings.Index(t, string(c))
	// 	if (indx == -1 ||  indx <= newIndex) {
	// 		return false
	// 	}
	// 	newIndex = indx
	// 	t = strings.Replace(t, string(t[indx]), "#", 1)
	// }
	// return true
	sLen := len(s)
	tLen := len(t)
	if sLen > tLen {
		return false
	}
	i := 0
	for j := 0; j < tLen && i < sLen; j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == sLen
}

func main() {
	s, t := "ab","bbba"
	fmt.Println(isSubsequence(s, t))
}
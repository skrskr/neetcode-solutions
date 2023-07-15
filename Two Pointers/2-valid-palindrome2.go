package main

import "fmt"

func isPalindrome(s string) bool {
	
	for i, j := 0, len(s) - 1;j > i;i,j = i+1,j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	sLen := len(s)

	for i, j := 0, sLen - 1; j > i; i, j = i+1, j-1 {
		if s[j] == s[i] {
			continue
		}
		return (isPalindrome(s[i+1: j+1]) || isPalindrome(s[i: j]))		
	}
	return true
}

func main() {
	s := "abc"
	fmt.Println(validPalindrome(s))
}
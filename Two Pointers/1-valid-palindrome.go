package main

import (
	"fmt"
)

func isAlphaNumber(ch byte) bool {
	
	if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') && (ch < '0' || ch > '9') {
		return false
	}
	return true
}

func isDigit(ch byte) bool {
	return (ch >= '0' && ch <= '9')
}

func isUpper(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z')
}

func toLower(ch byte) byte {
	if isDigit(ch) || ! isUpper(ch) {
		return ch
	}
	return byte(ch + 32)
}


func isPalindrome(s string) bool {
    sLen := len(s)
	i , j := 0, sLen - 1
	for ; j > i; {
		if ! isAlphaNumber(s[i]) {
			i++
			continue
		}
		if ! isAlphaNumber(s[j]) {
			j--
			continue
		}
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("raceacar"))

}
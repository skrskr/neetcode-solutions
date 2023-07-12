package main

import (
	"fmt"
	"strings"
)
func lengthOfLastWord(s string) int {
    // s = strings.Trim(s, " ")
	// words := strings.Split(s, " ")
	// return len(words[len(words)-1])

	sLen := len(s)
	count := 0
	for i:= sLen - 1; i >= 0; i-- {
		if count == 0 && s[i] == ' ' {
			continue
		}
		if s[i] == ' ' {
			break
		}
		count++
	}
	return count
}

func main() {
	s := "Hello Worl d   "
	fmt.Println(lengthOfLastWord(s))
}
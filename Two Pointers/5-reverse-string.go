package main

import (
	"fmt"
)

func reverseString(s []byte)  {
	//// Time complexity O(n), Space complexity O(1)
    i, j := 0, len(s)

	for ; j > i ; i ,j = i+1, j-1 {
		tmp := s[j]
		s[j] = s[i]
		s[i] = tmp
	}
	fmt.Println(s)

}

func main() {
	// var s []byte("hello")
	// reverseString(s)
}
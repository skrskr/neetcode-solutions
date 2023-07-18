package main

import (
	"fmt"
	"strings"
)


func mergeAlternately(word1 string, word2 string) string {
    // word1Len := len(word1)
    // word2Len := len(word2)
	// i, j := 0,0
	// str := ""
	// for ; i < word1Len && j <word2Len; i, j = i+1, j+1 {
	// 	str += string(word1[i]) + string(word2[j])
	// }

	// for ; i < word1Len; i++ {
	// 	str += string(word1[i])
	// }

	// for ; j < word2Len; j++ {
	// 	str += string(word2[j])
	// }
	// return str
	//// using strings.Builder more memory efficient
	word1Len := len(word1)
    word2Len := len(word2)
	i := 0
	var str strings.Builder
	for ; i < word1Len && i <word2Len; i= i+1 {
		str.WriteByte(word1[i])
		str.WriteByte(word2[i])
	}
	
	for ; i < word1Len; i++ {
		str.WriteByte(word1[i])
	}
	
	for ; i < word2Len; i++ {
		str.WriteByte(word2[i])
	}
	return str.String()
 }

func main() {
	fmt.Println(mergeAlternately("abcev", "pqr"))
}
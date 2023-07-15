package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
    ///// Brute force solution 
    // minStr := strs[0]
    // minLen := len(strs[0])
    // for _, val := range strs {
    //     if len(val) < minLen {
    //         minStr = val
    //         minLen = len(val)
    //     }
    // }
    // minPrefixLen := 1000
    // for _, val := range strs{
    //     i := 0
    //     for ; i < minLen ; i++ {
    //         if val[i] != minStr[i] {
    //             break
    //         }
    //     }
    //     // no common prefix
    //     if i == 0 {
    //         return ""
    //     }
    //     if i < minPrefixLen {
    //         minPrefixLen = i
    //     }
    // }
    // return minStr[:minPrefixLen]

    //// Better solution
    minStr := strs[0]
    res := ""
    for i := 0; i < len(minStr); i++ {
        for _, s := range strs {
            if i == len(s) || s[i] != minStr[i] {
                return res
            }
        }
        res += string(minStr[i])
    }
    return res
}


func main() {
	strs := []string{"dog","acecar","car"}
	fmt.Println(longestCommonPrefix(strs))
}
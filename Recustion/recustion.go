package main

import "fmt"


func printNums(n int) {
	if n < 0 {
		return 
	}
	
	// in desc order
	fmt.Println(n)

	printNums(n-1)
	// in asc order
	// fmt.Println(n)
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}



func main()  {
	// printNums(10)
	// fmt.Println(fact(5))
}
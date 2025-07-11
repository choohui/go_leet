package main

/*

 */

import (
	"fmt"
	"strconv"
)

func main() {
	input1 := "141"
	fmt.Println(isPalindrome(input1))
}

func isPalindrome(x int) bool {
	str_x := strconv.Itoa(x)
	for i, j := 0, len(str_x)-1; i <= j; i, j = i+1, j-1 {
		if str_x[i] != str_x[j] {
			return false
		}
	}
	return true
}

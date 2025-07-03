package main

/*
numbering each elements
2>check
find every possible combination
vaildate
 */

import (
	"fmt"
)

func main() {
	input1 := "babad"
	fmt.Println(longestPalindrome(input1))
}

func isitpalindrome(s_list []rune) bool {
	for i, j := 0, len(s_list)-1; i < j; i, j = i+1, j-1 {
		if s_list[i] != s_list[j] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	i_list := []rune(s)
	answer := ""
	for k := 0; k < len(i_list); k++ {
		for l := k; l < len(i_list); l++ {
			instance := i_list[k : l+1]
			if isitpalindrome(instance) {
				if len(instance) > len([]rune(answer)) {
					answer = string(instance)
				}
			}
		}
	}
	return answer
}
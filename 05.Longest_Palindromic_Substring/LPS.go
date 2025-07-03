/*
numbering each elements
2>check
find every possible combination
vaildate
 */

package main

import (
	"fmt"
)

func main() {
	input1 := "babad"
	fmt.Println(longestPalindrome(input1))
}

func findin(c rune, i_list []rune) bool {
	for _, v := range i_list {
		if v == c {
			return true
		}
	}
	return false
}

func longestPalindrome(s string) string {
	var i_list []rune
	for i := 0; i<len(s);i++ {
		i_list = append(i_list, s[i])
	}


}
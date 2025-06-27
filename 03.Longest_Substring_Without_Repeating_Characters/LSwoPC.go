/*
how to know chr exist in str
instant_list to store present flow

idea2 -> 중복되는
*/

package main

import (
	"fmt"
)

func main() {
	str_test := "abcdhgakdfk"
	fmt.Println(lengthOfLongestSubstring(str_test))
}

func findin(c rune, i_list []rune) bool {
	for _, v := range i_list {
		if v == c {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	if len(s) > 200 {
		s = s[:200]
	}
	max := 0
	len_s := len(s)
	instant_list := []rune{}
	for j := 0; j < len_s; j++ {
		for i := j; i < len_s; i++ {
			if !findin(rune(s[i]), instant_list) {
				instant_list = append(instant_list, rune(s[i]))
			} else {
				if max < len(instant_list) {
					max = len(instant_list)
				}
				instant_list = []rune{}
			}
		}
		if max < len(instant_list) {
			max = len(instant_list)
		}
		instant_list = []rune{}
	}

	return max
}

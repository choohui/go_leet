package main

/*

 */

import (
	"fmt"
)

func main() {
	input1 := "III"
	fmt.Println(romanToInt(input1))
}

func romanToInt(s string) int {
	result := 0
	vals := []int{1000, 500, 100, 50, 10, 5, 1}
	syms := []string{"M", "D", "C", "L", "X", "V", "I"}
	i_list := make([]int, len(s))
	for i := 0; i < len(syms); i++ {
		for j := 0; j < len(s); j++ {
			if syms[i] == string(s[j]) {
				i_list[j] = vals[i]
			}
		}
	}
	for k := 0; k < len(i_list); k++ {
		if k != len(i_list)-1 {
			if i_list[k+1] > i_list[k] {
				result += i_list[k+1] - i_list[k]
				k++
				continue
			}
		}
		result += i_list[k]
	}
	return result
}

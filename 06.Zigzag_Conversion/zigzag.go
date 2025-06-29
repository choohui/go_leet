package main

/*
	2n-2 pattern
	n row
*/

import (
	"fmt"
)

func main() {
	input1 := "PAYPALISHIRING"
	input2 := 3
	fmt.Println(convert(input1, input2))
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	pattern_n := numRows*2 - 2
	var where_row int
	var origin_list []rune = []rune(s)
	var result string
	var result_0 string
	if len(s) >= numRows {
		for k := 0; k < len(s); k++ {
			where_row = k % pattern_n
			if where_row == 0 {
				result_0 = result_0 + string(origin_list[k])
			}
		}
	}
	for i := 1; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			if j%pattern_n >= numRows {
				where_row = pattern_n - (j % pattern_n)
			} else {
				where_row = j % pattern_n
			}
			if where_row == i {
				result = result + string(origin_list[j])
			}
		}
		// fmt.Println(result) // 필요시 주석 해제
	}
	return result_0 + result
}

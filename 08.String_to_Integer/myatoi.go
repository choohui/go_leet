package main

import (
	"fmt"
)

func main() {
	input1 := "  -0012"
	fmt.Println(myAtoi(input1))
}



func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	s_array := []rune(s)
	var plus bool = true
	for k :=0; k<len(s_array); k++{
		if s_array[k] != ' ' {
			s_array = s_array[k:]
			break
		}
	}
	if s_array[0] == '+' {
		s_array = s_array[1:]
	} else if s_array[0] == '-' {
		s_array = s_array[1:]
		plus = false
	}
	for j :=0; j<len(s_array); j++{
		if s_array[j] != '0' {
			s_array = s_array[j:]
			break
		}
	}
	for i, v := range s_array {
		if v < '0' || v > '9' {
			s_array = s_array[:i]
			break
		}
	}
	result := 0
	const INT_MAX = 2147483647
	const INT_MIN = -2147483648
	for _, r := range s_array {
		digit := int(r - '0')
		if plus {
			if result > (INT_MAX-digit)/10 {
				return INT_MAX
			}
		} else {
			if result > (-(INT_MIN+digit))/10 {
				return INT_MIN
			}
		}
		result = result*10 + digit
	}
	if !plus {
		result = -result
	}
	return result
}
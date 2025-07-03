package main

/*

 */

import (
	"fmt"
	"strconv"
)

func main() {
	input1 := -1324525
	fmt.Println(reverse(input1))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverse(x int) int {
	if x < -2147483648 || x > 2147483647 {
		return 0
	}
	res := 0
	if x >= 0 {
		x_str := fmt.Sprintf("%d", x)
		x_str = reverseString(x_str)
		res, _ = strconv.Atoi(x_str)
	} else {
		x = -x
		x_str := fmt.Sprintf("%d", x)
		x_str = reverseString(x_str)
		res, _ = strconv.Atoi(x_str)
		res = -res
	}
	if res < -2147483648 || res > 2147483647 {
		return 0
	}
	return res
}

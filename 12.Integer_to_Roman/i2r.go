package main

/*

 */

import (
	"fmt"
)

func main() {
	input1 := 1355
	fmt.Println(intToRoman(input1))
}

func intToRoman(num int) string {
    vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    result := ""
    for i := 0; i < len(vals); i++ {
        for num >= vals[i] {
            num -= vals[i]
            result += syms[i]
        }
    }
    return result
}

/*

func intToRoman(num int) string {
	result := ""
	symbols := [][3]string{
		{"I", "V", "X"},
		{"X", "L", "C"},
		{"C", "D", "M"},
		{"M", "", ""}, 

	i := 0 
	for num != 0 {
		rest := num % 10
		first, second, third := symbols[i][0], symbols[i][1], symbols[i][2]
		instance := ""

		if rest == 4 {
			instance = first + second
		} else if rest == 9 {
			instance = first + third
		} else {
			if rest >= 5 {
				instance += second
				rest -= 5
			}
			for j := 0; j < rest; j++ {
				instance += first
			}
		}
		result = instance + result
		num /= 10
		i++
	}

	return result
}
	
*/

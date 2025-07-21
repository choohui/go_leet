package main

/*

 */

import (
	"fmt"
)

func main() {
	input1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(input1))
}

func longestCommonPrefix(strs []string) string {
	result := ""
	instance := ""
	for i := 0; i < len(strs[0]); i++ {
		instance = strs[0][i : i+1]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i {
				return result
			}
			if strs[j][i:i+1] != instance {
				return result
			}
		}
		result = result + instance
	}
	return result
}

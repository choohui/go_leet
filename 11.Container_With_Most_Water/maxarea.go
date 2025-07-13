package main

/*

 */

import (
	"fmt"
)

func main() {
	input1 := []int{1, 2, 5, 7}
	fmt.Println(maxArea(input1))
}

func maxArea(height []int) int {
	max_area := 0
	left := 0
	right := len(height) - 1
	for left < right {
		max_area = max(max_area, (right-left)*min(height[right], height[left]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max_area
}

// 	var width, length, area int
// 	for i := 0; i < len(height)-1; i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			length = j - i
// 			if height[i] > height[j] {
// 				width = height[j]
// 			} else {
// 				width = height[i]
// 			}
// 			area = length * width
// 			if max_area < area {
// 				max_area = area
// 			}

// 		}
// 	}
// 	return max_area
// }

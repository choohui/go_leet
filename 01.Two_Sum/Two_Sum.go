package main

import "fmt"

func findIndex(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		b_tar := target - v
		index := findIndex(nums[i+1:], b_tar)
		if index == -1 {
			continue
		} else {
			index = index + i + 1
			output := []int{i, index}
			return output
		}
	}
	return []int{0, 0}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

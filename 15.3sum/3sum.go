package main

/*

 */

import (
	"fmt"
	"sort"
)

func main() {
	input1 := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(input1))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicate values for i
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				triplet := []int{nums[i], nums[j], nums[k]}
				result = append(result, triplet)
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return result
}

/*
func isitin(a, b, c int, arr [][]int) bool {
	x := []int{a, b, c}
	sort.Ints(x)
	for _, triplet := range arr {
		y := make([]int, 3)
		copy(y, triplet)
		sort.Ints(y)
		if x[0] == y[0] && x[1] == y[1] && x[2] == y[2] {
			return true
		}
	}
	return false
}

func threeSum(nums []int) [][]int {
	len_n := len(nums)
	var result [][]int
	for i := 0; i < len_n-2; i++ {
		for j := i + 1; j < len_n-1; j++ {
			for k := j + 1; k < len_n; k++ {
				if isitin(nums[i], nums[j], nums[k], result) {
					break
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					fmt.Println(i, j, k)
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}
*/

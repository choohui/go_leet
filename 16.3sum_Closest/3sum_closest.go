package main

/*

 */

import (
	"fmt"
	"sort"
)

func main() {
	input1 := []int{-1, 0, 1, 2, -1, -4}
	input2 := 2
	fmt.Println(threeSumClosest(input1, input2))
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	instance := nums[0] + nums[1] + nums[2]
	inst_diff := absInt(instance - target)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			} else if sum < target {
				if target-sum < inst_diff {
					instance = sum
					inst_diff = target - sum
				}
				j++
			} else {
				if sum-target < inst_diff {
					instance = sum
					inst_diff = sum - target
				}
				k--
			}
		}
	}

	return instance
}

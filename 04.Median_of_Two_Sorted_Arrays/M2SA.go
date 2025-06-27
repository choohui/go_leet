/*

 */

package main

import (
	"fmt"
)

func main() {
	nums1 := []int{2, 3}
	nums2 := []int{1, 2, 3}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m_array := []int{}
	m_len := len(nums1) + len(nums2)
	n1i := 0
	n2i := 0
	for i := 0; i < m_len; i++ {
		if n1i >= len(nums1) {
			m_array = append(m_array, nums2[n2i])
			n2i++
		} else if n2i >= len(nums2) {
			m_array = append(m_array, nums1[n1i])
			n1i++
		} else if nums1[n1i] >= nums2[n2i] {
			m_array = append(m_array, nums2[n2i])
			n2i++
		} else {
			m_array = append(m_array, nums1[n1i])
			n1i++
		}
	}
	fmt.Println(m_array)
	if len(m_array)%2 == 1 {
		return float64(m_array[len(m_array)/2])
	} else {
		mid := len(m_array) / 2
		return float64(m_array[mid-1]+m_array[mid]) / 2.0
	}
}

// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays/description/


// Given two arrays of unique digits nums1 and nums2, return the smallest number that contains at least one digit from each array.

package problems

import (
	"fmt"
	"sort"
)

func Minimum_Number() {
	var nums1 = []int{3, 5, 2, 6}
	var num2 = []int{3, 1, 7}
	fmt.Println("Smallest number from two array = ", minNumber(nums1, num2))
}

func minNumber(nums1 []int, nums2 []int) int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	m := make(map[int]int, len(nums1))
	for _, value := range nums1 {
		m[value]++
	}
	for _, value := range nums2 {
		if m[value] != 0 {
			return value
		}
	}

	if nums1[0] == nums2[0] {
		return nums1[0]
	} else if nums1[0] > nums2[0] {
		return nums2[0]*10 + nums1[0]
	} else {
		return nums1[0]*10 + nums2[0]
	}
}

 
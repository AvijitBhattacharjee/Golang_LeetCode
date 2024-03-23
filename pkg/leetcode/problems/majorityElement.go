// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/majority-element/description/

// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

package problems

import "fmt"

func Majority_Element() {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(arr))
}

func majorityElement(nums []int) int {

	if len(nums) == 0 {
		return 0
	} else {
		fmt.Println("length = ", len(nums))
		var m = make(map[int]int)
		for _, value := range nums {
			m[value]++
		}
		fmt.Println(m)
		for index, value := range m {
			if value > (len(nums) / 2) {
				return index
			}
		}
	}
	return 0
}

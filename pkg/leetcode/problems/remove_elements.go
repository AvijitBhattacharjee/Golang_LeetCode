// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/remove-element/description/

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
// Return k.



package problems

import "fmt"

func Remove_element() {

	var arr = []int{0, 1, 2, 2, 3, 0, 4, 2}
	var val = 2
	fmt.Println("Solving remove element from an array")
	fmt.Println(removeelement(arr, val))
}
func removeelement(arr []int, val int) int {

	var j int
	for i := 0; i < len(arr); i++ {
		if arr[i] != val {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
	fmt.Println("printing the array", arr)
	return j

}

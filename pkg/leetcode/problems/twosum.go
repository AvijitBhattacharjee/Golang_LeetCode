// Copyright (c) avijit bhattacharjee 2024

package problems

import "fmt"

func Twosum_solution() {
	fmt.Println("Solving LeetCode problem 1...")
	// Add your solution for LeetCode problem 1 here

	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(twoSum(arr, 6))
}

func twoSum(nums []int, target int) []int {

	var flag1, flag2, a int
	var flag = false
	for index, value := range nums {
		a = target - value
		for index2, value2 := range nums {
			if a == value2 && index2 != index {
				flag2 = index2
				flag1 = index
				flag = true
				break
			}

		}
		if flag == true {
			break
		}
	}
	a1 := []int{flag1, flag2}
	return a1

}

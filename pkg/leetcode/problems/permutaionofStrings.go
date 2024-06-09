// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/permutation-difference-between-two-strings/

package problems

import (
	"fmt"
)

func FindPermutationDifference() {

	var s = "abc"
	var t = "bac"
	fmt.Println(findPermutationDifference(s, t))
}

func findPermutationDifference(s string, t string) int {

	var m1 = make(map[rune]int)
    var m2 = make(map[rune]int)
	var result int

    for index, value := range s {
        m1[value] = index
    }

    for index, value := range t {
        m2[value] = index
    }

	for index, value := range m1 {
		result = result + absDiff(m2[index], value)
	}
	return (result)
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

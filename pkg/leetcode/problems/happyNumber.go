// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/happy-number/


package problems

import (
	"fmt"
)

func HappyNumber() {
	n := 19
	fmt.Println(isHappy(n))
}

func isHappy(n int) bool {

	// using map because of the cycle, if it is not happy number after n number of iteration n will again come and the map will throw them false
	seen := map[int]bool{}
	for n != 1 && !seen[n] {
		seen[n] = true
		n = sumOfSqaureOfDigits(n)
	}
	return n == 1
}

func sumOfSqaureOfDigits(n int) int {

	res := 0
	for n > 0 {
		res = res + (n%10)*(n%10)
		n = n / 10
	}
	fmt.Println(res)
	return res
}
// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/ugly-number/description/

// An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
// Given an integer n, return true if n is an ugly number.

package problems

import "fmt"

func UglyNumber() {
	var num = 9
	fmt.Println(isUgly(num))
}

func isUgly(n int) bool {

	if n < 0 {
		return false
	}
	if n == 1 {
		return true
	}
	m := make(map[int]int)
	for n%2 == 0 {
		n = n / 2
		m[2]++
	}
	fmt.Println(n)

	for i := 3; i*i <= n; i = i + 2 {
		fmt.Println("number is = ", n)
		for n%i == 0 {
			n = n / i
			m[i]++
		}
	}

	if n > 2 {
		m[n]++
	}

	fmt.Println(m)
	for index, _ := range m {
		if index > 5 {
			return false
		}
	}
	return true
}

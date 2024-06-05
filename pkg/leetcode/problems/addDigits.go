// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/add-digits/


package problems

import (
	"fmt"
)

func AddDigits() {
	n := 0
	fmt.Println(addDigits(n))
}

func addDigits(num int) int {
    
	if sumOfDigits(num) < 10 {
		return sumOfDigits(num)
	}
	return addDigits(sumOfDigits(num))
}


func sumOfDigits(num int) int {

	var res = 0
	for(num>0) {
		res = res + num%10
		num = num/10
	}
	return res
}
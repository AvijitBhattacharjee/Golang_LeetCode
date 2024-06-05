// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/

package problems

import (
	"fmt"
)

func GetLucky() {
	var str = "zbax"
	var k = 0
	fmt.Println(getLucky(str, k))
}

func getLucky(s string, k int) int {
    sum := 0
    for _, val := range s {
        value := int(val - 'a' + 1)
	    sum += value % 10 + value / 10
    }

	for k > 1 {
		newSum := 0
		for sum > 0 {
			newSum += sum%10
			sum /= 10
		}
		sum = newSum
		k--
	}
    return sum
}
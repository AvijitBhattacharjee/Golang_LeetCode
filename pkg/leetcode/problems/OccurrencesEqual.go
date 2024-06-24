// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/


package problems

import (
	"fmt"
	//"strings"
)

func OccurrencesEqual() {
	fmt.Println(areOccurrencesEqual("abacbc"))
} 

func areOccurrencesEqual(s string) bool {
    
	var m = make(map[rune]int)
	for _, value := range s{
		m[value]++
	}
	return mapAllValueSame(m)
}


func mapAllValueSame(m map[rune]int) bool {

	var flag, firstValue = false,0
	
	for _, value := range m {
		if !flag {
			firstValue = value
			flag = true
		} else {
			if value != firstValue {
				return false
			}
		}
	}
	return true
}
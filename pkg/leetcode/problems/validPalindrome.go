// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/valid-palindrome/description/


package problems

import (
	"fmt"
	"unicode"
)

func Palindrome() {
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {

	var validString string
	for _, value := range s {
		if value >= 48 && value <= 57 {
			validString = validString + string(value)
		}
		if unicode.IsLetter(value) {
			if value >= 65 && value <= 90 {
				validString = validString + string(value+32)		
			} else {
				validString = validString + string(value)
			}
		}
	}
	fmt.Println(validString)
	for i:=0; i<len(validString); i++ {
		if validString[i] != validString[len(validString)-1-i] {
			return false
		}
	}
 
	return true
}
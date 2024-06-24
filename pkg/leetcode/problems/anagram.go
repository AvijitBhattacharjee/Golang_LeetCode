// Copyright (c) avijit bhattacharjee 2024

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

package problems

import "fmt"

func Anagram() {

	var str1 = "avijit"
	var str2 = "jitavi"

	fmt.Println("Solving Anagram")
	fmt.Println(compare(str1, str2))
}


func compare(str1, str2 string) bool {
	return getKey(str1) == getKey(str2)
}

func getKey(str string) string {
	// Normalize the string to lowercase to handle case insensitivity
	str = strings.ToLower(str)

	// Create a slice to count character frequencies
	var res = make([]int, 26)

	for _, value := range str {
		if value >= 'a' && value <= 'z' { // Ensure only lowercase letters are processed
			res[value-'a']++
		}
	}
	return fmt.Sprint(res)
}
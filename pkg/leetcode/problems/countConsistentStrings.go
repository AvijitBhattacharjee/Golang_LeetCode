// Copyright (c) avijit bhattacharjee 2024


// https://leetcode.com/problems/count-the-number-of-consistent-strings/description/


package problems

import "fmt"

func CountingConsistentStrings() {

	var allowed = "abc"
	var words = []string{"a","b","c","ab","ac","bc","abc"}
	fmt.Println(countConsistentStrings(allowed, words))
}

func countConsistentStrings(allowed string, words []string) int {
    

	var myMap = make(map[rune]int, len(allowed))
	
	for _, value := range allowed {
		myMap[value]++
	}

	fmt.Println(myMap)
	var count int


	for _, word := range words {

		for _, char := range word {
			if myMap[char] != 1 {
				count++
				break
			}
		}

	}


	return (len(words) - count)



}

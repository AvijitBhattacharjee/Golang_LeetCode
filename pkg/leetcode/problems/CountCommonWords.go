// Copyright (c) avijit bhattacharjee 2024


// https://leetcode.com/problems/count-common-words-with-one-occurrence/description/


package problems

import "fmt"

func CountCommonWords() {
	var words1 = []string{"leetcode","is","amazing","as","is"}
	var words2 = []string{"amazaing","leetzcode","ais"}
	fmt.Println(countWords(words1, words2))

}

func countWords(words1 []string, words2 []string) int {
    
	var map1, map2 = make(map[string]int), make(map[string]int)

	for _, value := range words1 {
		map1[value]++
	}
	
	for _, value := range words2 {
		map2[value]++
	}

	var count int

	for index, value := range map1 {
		if value == 1 && map2[index] == 1 {
			count++
		}
	}

	return count


}
// Copyright (c) avijit bhattacharjee 2024


// https://leetcode.com/problems/word-break/description/


package problems

import "fmt"

func WordBreak() {
	var s = "catsandog"
	var str = []string{"cats","dog","sand","and","cat"}
	fmt.Println(wordBreak(s, str))
}

func wordBreak(s string, wordDict []string) bool {
    
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i<= len(s) ; i++ {
		for j := i-1; j>= max(i - max_length(wordDict) - 1, 0); j-- {
			if dp[j] && contains(wordDict, s[j:i]) {
				dp[i] = true
				break
			}
		}
	}	
	return dp[len(s)]

}


func contains(words []string, word string) bool {

	for _, value := range words {
		if value == word {
			return true
		}
	}
	return false
}

func max_length(words []string) int{
	
	var maximum_lenth int
	for _, word := range words {
		if maximum_lenth < len(word) {
			maximum_lenth = len(word)
		}
	}
	return maximum_lenth
}

func max(a, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}
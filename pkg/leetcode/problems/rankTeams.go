// Copyright (c) avijit bhattacharjee 2024


// https://leetcode.com/problems/rank-teams-by-votes/


package problems

import (
	"fmt"
	"sort"
)


func RankTeams() {
	var votes = []string{"WXYZ","XYZW"}
	fmt.Println(rankTeams(votes))
	
}



func rankTeams(votes []string) string {

	var result = make(map[rune]int, len(votes[0]))


	for _, vote := range votes {

		for index, value := range vote {
			result[value] = result[value] + index
		}
	}

	var pairs []struct{ key, value int }
	for k, v := range result {
		pairs = append(pairs, struct{ key, value int }{k, v})
	}

	// Sort the slice based on the values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	// Create a new map from the sorted slice
	sortedMap := make(map[int]int)
	for _, p := range pairs {
		sortedMap[p.key] = p.value
	}

	fmt.Println(sortedMap)
	return "hahaha"
    
}
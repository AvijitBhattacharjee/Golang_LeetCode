// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/rank-teams-by-votes/

package problems

import (
	"fmt"
	"sort"
)

func RankTeams() {
	var votes = []string{"WXYZ", "XYZW"}
	fmt.Println(rankTeams(votes))

}

func rankTeams(votes []string) string {

	var result = make(map[rune]int, len(votes[0]))

	for _, vote := range votes {

		for index, value := range vote {
			result[value] = result[value] + index
		}
	}
	fmt.Println(result)

	keys := make([]rune, 0, len(result))

	for key := range result {
		keys = append(keys, key)
	}

	fmt.Println(result)
	fmt.Println(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return result[keys[i]] < result[keys[j]]
	})

	fmt.Println(keys)
	var res string

	for _, value := range keys {
		res = res + string(value)
	}

	fmt.Println(res)

	return "hahaha"

}

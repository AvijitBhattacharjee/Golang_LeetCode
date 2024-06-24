// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/sort-characters-by-frequency/

package problems

import (
	"fmt"
	"sort"
)

func FrequencyStringSort() {
	fmt.Println(frequencySort("tree"))
}

func frequencySort(s string) string {
    var freq = make([]int, 255)

    for i:=0;i<len(s);i++ {
        freq[s[i]]++
    }

    res := []byte(s)

    sort.Slice(res, func(i,j int) bool  {
        if freq[res[i]] == freq[res[j]] {
            return res[i] > res[j]
        }
        return freq[res[i]] > freq[res[j]]
    })
    return string(res)
}
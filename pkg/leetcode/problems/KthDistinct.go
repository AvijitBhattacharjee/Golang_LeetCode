// Copyright (c) avijit bhattacharjee 2024


// https://leetcode.com/problems/kth-distinct-string-in-an-array/description/


package problems

import "fmt"

func KthDistinct() {

	var arr = []string{"d","b","c","b","c","a"}
	var k = 2
	fmt.Println(kthDistinct(arr, k))

}

func kthDistinct(arr []string, k int) string {

	var output, res = "", make(map[string]int)

	for _, value := range arr {
		res[value]++
	}
	
	for _, value := range arr {
		if res[value] == 1 {
			k--
		}
		if k == 0 {
			return value
		}
	}
	return output
    
}
// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/find-the-difference/description/

package problems

import "fmt"

func Practice() {

	var s = "abcd"
	var t = "abcde"
	fmt.Println(findTheDifference(s, t))
}

func findTheDifference(s string, t string) byte {
    
    var m1 = make(map[rune]int)
    var m2 = make(map[rune]int)

    for _, value := range s {
        m1[value]++
    }

    for _, value := range t {
        m2[value]++
    }

    fmt.Println(m1)
	fmt.Println(m2)

	for index, _ := range m2 {
		if m2[index] != m1[index] {
			return byte(index)
		}
	}

	return byte(0)
}
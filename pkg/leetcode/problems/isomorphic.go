// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/isomorphic-strings/description/


// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

package problems

import (
	"fmt"
)

func Isphormic() {
	var str1 = "badc"
	var str2 = "baba"
	fmt.Println(isIsomorphic(str1, str2))
}
func isIsomorphic(s string, t string) bool {

	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		if m[int(s[i])] == 0 {
			m[int(s[i])] = int(t[i])
		} else {
			if m[int(s[i])] != int(t[i]) {
				return false
			}
		}
	}
	n := make(map[int]int)

	for i := 0; i < len(s); i++ {
		if n[int(t[i])] == 0 {
			n[int(t[i])] = int(s[i])
		} else {
			if n[int(t[i])] != int(s[i]) {
				return false
			}
		}
	}
	return true
}


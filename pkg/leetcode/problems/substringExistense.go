// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/


package problems

import (
	"fmt"
	"strings"
)

func SubstringPresent() {
	fmt.Println(isSubstringPresent("leetcode"))
}


func isSubstringPresent(s string) bool {
	for i:=len(s)-1; i>0; i-- {	if strings.Contains(s, string(s[i])+string(s[i-1])) {return true} }
	return false
}

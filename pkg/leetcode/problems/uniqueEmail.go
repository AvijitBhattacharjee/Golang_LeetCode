// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/unique-email-addresses/description/

package problems

import (
	"fmt"
	"strings"
)


func UniqueEmail() {

	var emails = []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com", "b@leetcode.com","c@leetcode.com"}
	fmt.Println(numUniqueEmails(emails))
}

func numUniqueEmails(emails []string) int {

	var emailMap = make(map[string]int)
	for _, email := range emails {
		
		domainName := strings.Split(email, "@")[1]
		var username string
		userNameWithoutDot := strings.Split(strings.Split(strings.Split(email, "@")[0], "+")[0], ".")

		for _, value := range userNameWithoutDot {
			username = username + value
		}

		validEmail := username + "@" + domainName
		emailMap[validEmail]++
	}
	return len(emailMap)
}

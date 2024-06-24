// Copyright (c) avijit bhattacharjee 2024

// Given an array of logs with the server name and status (either failed or successful) had to return the number of servers that failed 3 times in a row.


package problems

import (
	"fmt"
	"strings"
)

func ServerFail() {

	logs := []string{"serverA failed", "serverA failed", "serverB failed", "serverB failed", "serverA failed", "serverB failed"}
	fmt.Println(serverFailCount(logs))
}

func serverFailCount(logs []string) int {
	var result = make(map[string][]string)

	for _, value := range logs {
		contents := strings.Split(value, " ")
		result[contents[0]] = append(result[contents[0]], contents[1])
	}
	var res []string

	for index, value := range result {
		if consequitive(value) {
			res = append(res, index)
		}
	}
	return len(res)
}

func consequitive(str []string) bool {

	for i:=0;i<len(str)-2;i++ {
		if str[i] == "failed" && str[i+1] == "failed" && str[i+2] == "failed" {
			return true
		}
	}
	return false
}
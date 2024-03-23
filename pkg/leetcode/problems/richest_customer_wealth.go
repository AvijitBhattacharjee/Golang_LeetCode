// Copyright (c) avijit bhattacharjee 2024

// https://leetcode.com/problems/richest-customer-wealth/description/


// You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.
//A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.

package problems

import "fmt"

func Max_Wealth() {
	var wealth = [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Println("calculating richest wealth = ", maximumWealth(wealth))
}

func maximumWealth(accounts [][]int) int {

	if len(accounts) == 0 {
		return 0
	}
	max := 0
	for _, value := range accounts {
		sum := 0
		for _, value2 := range value {
			sum = sum + value2
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// Copyright (c) avijit bhattacharjee 2024

package problems

import (
	"fmt"
	"math/big"
)

func NumTrees() {
	fmt.Println(numTrees(11))
}

func numTrees(n int) int {
	nBig := big.NewInt(int64(n))
	twoNBig := new(big.Int).Mul(big.NewInt(2), nBig)
	nPlusOneBig := new(big.Int).Add(nBig, big.NewInt(1))

	num := factorial(twoNBig)
	den1 := factorial(nBig)
	den2 := factorial(nPlusOneBig)

	den := new(big.Int).Mul(den1, den2)
	result := new(big.Int).Div(num, den)
	return int(result.Int64())

}

func factorial(n *big.Int) *big.Int {
	// Base case: 0! = 1 and 1! = 1
	one := big.NewInt(1)
	if n.Cmp(one) <= 0 {
		return one
	}

	// n! = n * (n-1)!
	nMinusOne := new(big.Int).Sub(n, one)
	return new(big.Int).Mul(n, factorial(nMinusOne))
}

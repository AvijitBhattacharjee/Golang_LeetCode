// Copyright (c) avijit bhattacharjee 2024

package problems

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rangeSumBST(low int, high int) int {

    var res *int
	res 
    // if root == nil {
    //     return res
    // }
    // if root.Val >= low && root.Val <= high {
    //     res = res + root.Val
    // }
    // rangeSumBST(root.Left, low, high)
    // rangeSumBST(root.Right, low, high)
	if res >= low && res <= high {
		return res
	}
    return 0
}

func main() {
	fmt.Println(rangeSumBST())
}
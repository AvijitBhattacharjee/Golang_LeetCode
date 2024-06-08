// Copyright (c) avijit bhattacharjee 2024

package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  func rightSideView(root *TreeNode) []int {
//     result := []int{}
//     rightView(root, &result, 0)
//     return result
// }

// // rightView helper function to traverse the tree
// func rightView(curr *TreeNode, result *[]int, currDepth int) {
//     if curr == nil {
//         return
//     }
//     if currDepth == len(*result) {
//         *result = append(*result, curr.Val)
//     }

//     rightView(curr.Right, result, currDepth+1)
//     rightView(curr.Left, result, currDepth+1)
// }

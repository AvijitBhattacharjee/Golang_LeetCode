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

//  func insertIntoBST(root *TreeNode, val int) *TreeNode {

//     if (root == nil) {
//         return &TreeNode{Val: val}
//     }
//     if (root.Val > val) {
//         root.Left = insertIntoBST(root.Left, val)
//     }
//     if (root.Val < val) {
//         root.Right = insertIntoBST(root.Right, val)
//     }
//     return root
// }
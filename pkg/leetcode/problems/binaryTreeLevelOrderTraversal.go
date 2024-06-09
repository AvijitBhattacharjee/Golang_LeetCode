// Copyright (c) avijit bhattacharjee 2024

package problems

// https://leetcode.com/problems/binary-tree-level-order-traversal/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  func levelOrder(root *TreeNode) [][]int {
    
//     var m = make(map[int][]int)
//     getArray(root, 0, m)
//     s := make([][]int, len(m))
//     for index, value := range m {
//         s[index] = value
//     }
//     return s
// }

// func getArray(root *TreeNode, level int, m map[int][]int) {

//     if root == nil {
//         return
//     }

//     m[level] = append(m[level], root.Val)
//     getArray(root.Left, level+1, m)
//     getArray(root.Right, level+1, m)
// }
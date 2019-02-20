package leetCode

/*
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		s     = []*TreeNode{root}
		resp  = make([][]int, 0)
		bLToR = true
	)

	for len(s) > 0 {
		var (
			n   = len(s)
			tmp = make([]int, n)
		)

		if !bLToR { //right->left
			for i := 0; i < n; i++ {
				tmp[n-i-1] = s[i].Val

				if s[i].Left != nil {
					s = append(s, s[i].Left)
				}
				if s[i].Right != nil {
					s = append(s, s[i].Right)
				}

			}
		} else { //left->right
			for i := 0; i < n; i++ {
				tmp[i] = s[i].Val

				if s[i].Left != nil {
					s = append(s, s[i].Left)
				}
				if s[i].Right != nil {
					s = append(s, s[i].Right)
				}
			}
		}
		s, bLToR = s[n:len(s)], !bLToR
		resp = append(resp, tmp)
	}

	return resp
}

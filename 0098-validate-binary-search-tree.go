package leetCode

/*
https://leetcode.com/problems/validate-binary-search-tree/
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
Example 1:

Input:
    2
   / \
  1   3
Output: true
Example 2:

    5
   / \
  1   4
     / \
    3   6
Output: false
Explanation: The input is: [5,1,4,null,null,3,6]. The root node's value
             is 5 but its right child's value is 4.
*/

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//M1: 采用从上往下比较的方法：所有的左子树小于根节点，所有的右子树节点小于根节点
const (
	INT_MAX = int(^uint(0) >> 1)
	INT_MIN = ^INT_MAX
)

func IsValidBST(root *TreeNode) bool {
	min, max := INT_MIN, INT_MAX
	return isValidBSTProxy(root, &min, &max)
}

func isValidBSTProxy(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}

	if *min != INT_MIN && root.Val <= *min {
		return false
	}
	if *max != INT_MAX && root.Val >= *max {
		return false
	}

	return isValidBSTProxy(root.Left, min, &root.Val) &&
		isValidBSTProxy(root.Right, &root.Val, max)
}

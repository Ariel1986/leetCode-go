package leetCode

/*
https://leetcode.com/problems/binary-tree-inorder-traversal/

Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
*/

/**
 * Definition for a binary tree node.
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//M1: 递归
func inorderTraversalProxy(root *TreeNode, resp *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorderTraversalProxy(root.Left, resp)
	}
	*resp = append(*resp, root.Val)

	if root.Right != nil {
		inorderTraversalProxy(root.Right, resp)
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	inorderTraversalProxy(root, &result)

	return result
}

//M2: Stack
class Solution2 {
   public:
       std::vector<int> inorderTraversal(TreeNode* root) {
           if(root == nullptr){
               return {};
           }
           
           std::vector<int> result;
           std::stack<TreeNode*> s;
           TreeNode *cur = root;
           while (cur != nullptr || !s.empty()) {
               if(cur != nullptr){
                   s.push(cur);
                   cur = cur->left;
               }else{
                   cur = s.top();
                   s.pop();
                   result.emplace_back(cur->val);
                   cur = cur->right;
               }
           }
           
           return result;
       }
   };
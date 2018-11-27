#include <algorithm>
#include <functional>
#include <vector>
/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]

        _______3______
       /              \
    ___5__          ___1__
   /      \        /      \
   6      _2       0       8
         /  \
         7   4
Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
Example 2:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself
             according to the LCA definition.
Note:

All of the nodes' values will be unique.
p and q are different and both values will exist in the binary tree.

*/

/*
Method 1: 递归
如果在root的左右子树中分别找到p, q,则root就是LCA节点
如果左子树中没有找到p和q，则p，q在root的右子树中，right就是LCA节点
同理，left是LCA节点

Method 2:PATH
从根->p, 根->q 的path中，找到最后一个相等的节点,就是结果
*/

/**
 * Definition for a binary tree node.
 * */
  struct TreeNode {
      int val;
      TreeNode *left;
      TreeNode *right;
      TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  };
 
 //M1:递归
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root == nullptr || root == p || root == q){
			return root;
		}
    
		TreeNode* left = lowestCommonAncestor(root->left, p, q);
		TreeNode *right = lowestCommonAncestor(root->right, p, q);

		if(left == nullptr){ return right; }
		if(right == nullptr){ return left; }
		return root;
		//return left == nullptr? right : right == nullptr? left : root;
	}
};

//M2: PATH
class Solution {
public:
    bool findPath(TreeNode *root, TreeNode *target, std::vector<TreeNode*>& path){
        if(root == nullptr){ return false; }
        if(root == target){
            path.emplace_back(root);
            return true;
        }
        path.emplace_back(root);
        if(findPath(root->left, target, path) || findPath(root->right, target, path)){
            return true;
        }
        
        path.pop_back();
        return false;
    }
    
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root == nullptr || root == p || root == q){
            return root;
        }
        
        std::vector<TreeNode*> pPath;
        std::vector<TreeNode*> qPath;
        findPath(root, p, pPath);
        findPath(root, q, qPath);
        
        TreeNode* resp = nullptr;
        
        for(int i = 0, iEnd = std::min(pPath.size(), qPath.size()); i < iEnd; i++){
            if(pPath[i] == qPath[i]){
                resp = pPath[i];
            }else{
                break;
            }
        }
        
        return resp;
    }
};
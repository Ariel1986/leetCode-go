#include <algorithm>
#include <functional>

/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given binary search tree:  root = [6,2,8,0,4,7,9,null,null,3,5]

        _______6______
       /              \
    ___2__          ___8__
   /      \        /      \
   0      _4       7       9
         /  \
         3   5
Example 1:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
Example 2:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself 
             according to the LCA definition.
Note:

All of the nodes' values will be unique.
p and q are different and both values will exist in the BST.
*/

/*
M1: 递归
M2: 循环
*/

/**
 * Definition for a binary tree node.
*/

//M1
struct TreeNode {
      int val;
      TreeNode *left;
      TreeNode *right;
      TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };

class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(p->val > q->val){ std::swap(p, q); }
        
        std::function<TreeNode*(TreeNode* root, TreeNode* p, TreeNode* q)> lowestCommonAncestorHelper =
        [&](TreeNode* root, TreeNode* p, TreeNode* q)->TreeNode*{
            if(root == nullptr){ return nullptr; }
            if(root->val < p->val){
                return lowestCommonAncestorHelper(root->right, p, q);
            }
            if(root->val > q->val){
                return lowestCommonAncestorHelper(root->left, p, q);
            }
            return root;
        };
        
        return lowestCommonAncestorHelper(root, p, q);
    }
};

//M2:
class Solution2 {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(p->val > q->val){ std::swap(p, q); }
        
        TreeNode *result = root;
        while(result != nullptr){
            if(result == p || result == q ||
               (result->val > p->val && result->val < q->val)){
                return result;
            }
            
            if(result->val > q->val){
                result = result->left;
            }else {
                result = result->right;
            }
        }
        
        return result;
    }
};
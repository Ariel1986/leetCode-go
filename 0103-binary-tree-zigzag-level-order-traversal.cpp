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
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    std::vector<std::vector<int>> zigzagLevelOrder(TreeNode* root) {
        if(root == nullptr){
            return {};
        }
        
        std::vector<int> vec;
        std::vector<std::vector<int>> result;
        std::queue<TreeNode*> q;
        bool LToR = true;
        
        q.emplace(root);
        q.emplace(nullptr);
        while(!q.empty()){
            auto p = q.front();
            q.pop();
            
            if(p == nullptr ){
                if(!q.empty()){ q.emplace(nullptr); }
                if(!LToR){
                    std::reverse(vec.begin(), vec.end());
                }
                LToR = !LToR;
                result.emplace_back(vec);
                vec.clear();
            }else{
                vec.emplace_back(p->val);
                if(p->left != nullptr){ q.emplace(p->left); }
                if(p->right != nullptr){ q.emplace(p->right); }
            }
        }
        return result;
    }
};
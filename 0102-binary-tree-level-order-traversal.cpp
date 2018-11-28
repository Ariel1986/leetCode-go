/*
https://leetcode.com/problems/binary-tree-level-order-traversal/

Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]

*/

class Solution {
public:
    std::vector<std::vector<int>> levelOrder(TreeNode* root) {
        if(root == nullptr){
            return {};
        }
        
        std::vector<int> vec;
        std::vector<std::vector<int>> result;
        std::queue<TreeNode*> q;
        
        q.emplace(root);
        q.emplace(nullptr);
        while(!q.empty()){
            auto p = q.front();
            q.pop();
            if(p == nullptr ){
                if(!q.empty()){ q.emplace(nullptr); }
                result.emplace_back(vec);
                vec = {};
            }else{
                vec.emplace_back(p->val);
                if(p->left != nullptr){ q.emplace(p->left); }
                if(p->right != nullptr){ q.emplace(p->right); }
            }
        }
        return result;
    }
};

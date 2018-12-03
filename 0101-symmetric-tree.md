<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/symmetric-tree/

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3
Note:
Bonus points if you could solve it both recursively and iteratively.
```

- [Recursion](#recursion)
- [Queue](#queue)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Recursion

```cpp
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
    bool isSymmetric(TreeNode* root) {
        if(root == nullptr){ return true; }
        std::function<bool(TreeNode* left, TreeNode* right)>
        isSymmetricProxy = [&](TreeNode* left, TreeNode* right)->bool{
            if(left == nullptr || right == nullptr){
                return left == right;
                
            }
            if(left->val != right->val){
                return false;
            }
            
            return isSymmetricProxy(left->right, right->left) && isSymmetricProxy(left->left, right->right);
        };
        
        return isSymmetricProxy(root->left, root->right);
    }
};

```

# Queue

```cpp
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        if(root == nullptr){ return true; }
        std::queue<TreeNode*> q;
        q.emplace(root->left);
        q.emplace(root->right);
        while(!q.empty()){
            auto left = q.front();
            q.pop();
            auto right = q.front();
            q.pop();
            
            if(left == nullptr && right == nullptr){
                continue;
            }
            if(left == nullptr ^ right == nullptr){
                return false;
            }
            if(left->val != right->val){
                return false;
            }
            q.emplace(left->left);
            q.emplace(right->right);
            q.emplace(left->right);
            q.emplace(right->left);
        }
        return true;
    }
};

```

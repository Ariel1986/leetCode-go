<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/invert-binary-tree/

Invert a binary tree.

Example:

Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1
Trivia:
This problem was inspired by this original tweet by Max Howell:

Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.
```

- [Recursion](#recursion)
- [Stack](#stack)

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
    TreeNode* invertTree(TreeNode* root) {
        if(root == nullptr){
            return root;
        }
        std::swap(root->left, root->right);
        invertTree(root->left);
        invertTree(root->right);
        return root;
    }
};

```

# Stack

```cpp
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        if(root == nullptr){
            return root;
        }
        std::stack<TreeNode*> s;
        s.emplace(root);
        while(!s.empty()){
            auto tmp = s.top();
            s.pop();
            std::swap(tmp->left, tmp->right);
            if(tmp->left != nullptr){
                s.emplace(tmp->left);
            }
            if(tmp->right != nullptr){
                s.emplace(tmp->right);
            }
        }
        
        return root;
    }
};
```
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/diameter-of-binary-tree/

Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

Example:
Given a binary tree 
          1
         / \
        2   3
       / \     
      4   5    
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.

```

- [Recursion](#recursion)

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
    int diameterOfBinaryTree(TreeNode* root) {
        if(!root) { return  0; }
        std::function<int(TreeNode* , int&)>
        depthOfTree = [&](TreeNode* root, int& diameter)->int{
            if(!root){ return 0; }
            auto lh = depthOfTree(root->left, diameter);
            auto rh = depthOfTree(root->right, diameter);
            diameter = std::max(diameter, lh + rh);
            return 1 + std::max(lh, rh);
        };
        
        int diameter = 0;
        depthOfTree(root, diameter);
        
        return diameter;
    }
};
```

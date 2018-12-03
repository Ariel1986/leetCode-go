<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/merge-two-binary-trees/

Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Example 1:

Input: 
	Tree 1                     Tree 2                  
          1                         2                             
         / \                       / \                            
        3   2                     1   3                        
       /                           \   \                      
      5                             4   7                  
Output: 
Merged tree:
	     3
	    / \
	   4   5
	  / \   \ 
	 5   4   7

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
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if(!t1 && !t2){ return nullptr; }
        if(!t1){ return t2; }
        if(!t2){ return t1; }
        
        TreeNode* result = nullptr;
        result = new TreeNode(t1->val + t2->val);
        result->left = mergeTrees(t1->left, t2->left);
        result->right = mergeTrees(t1->right, t2->right);
        
        return result;
    }
};
```

```cpp
class Solution2 {
public:
    TreeNode* mergeTrees(TreeNode* t1, TreeNode* t2) {
        if(!t1 && !t2){ return nullptr; }

        TreeNode* result = nullptr;
        if(!t1){
            result = new TreeNode(t2->val);
            result->left = mergeTrees(nullptr, t2->left);
            result->right = mergeTrees(nullptr, t2->right);
        }else{
            if(!t2){
                result = new TreeNode(t1->val);
                result->left = mergeTrees(t1->left, nullptr);
                result->right = mergeTrees(t1->right, nullptr);
            }
            else{
                result = new TreeNode(t1->val + t2->val);
                result->left = mergeTrees(t1->left, t2->left);
                result->right = mergeTrees(t1->right, t2->right);
            }
        }
     
        return result;
    }
};

```

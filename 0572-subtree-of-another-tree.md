<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/subtree-of-another-tree/

Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.

Example 1:
Given tree s:

     3
    / \
   4   5
  / \
 1   2
Given tree t:
   4 
  / \
 1   2
Return true, because t has the same structure and node values with a subtree of s.
Example 2:
Given tree s:

     3
    / \
   4   5
  / \
 1   2
    /
   0
Given tree t:
   4
  / \
 1   2
Return false.
```

# Hints
```
1. Which approach is better here - recursive or iterative?
2. If recursive approach is better, can you write recursive function with its parameters?
3. Two trees s and t are said to be identical if their root values are same and their left and right subtrees are identical. Can you write this in form of recursive formulae?
4. Recursive formulae can be: isIdentical(s,t)= s.val==t.val AND isIdentical(s.left,t.left) AND isIdentical(s.right,t.right)
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
    bool isSubtree(TreeNode* s, TreeNode* t) {
        if(!s){ return t == s; }
        if(!t){ return true; }
        
        std::function<bool(TreeNode* s, TreeNode* t)>
        isSubtreeProxy = [&](TreeNode* s, TreeNode* t) ->bool{
            if(!s){ return s == t; }
            if(!t){ return false; }
            if(s->val != t->val){ return false ;}
            return isSubtreeProxy(s->left, t->left) && isSubtreeProxy(s->right, t->right);
        };
        
        std::stack<TreeNode*> sta;
        sta.emplace(s);
        while(!sta.empty()){
            auto tmp = sta.top();
            sta.pop();
            if(isSubtreeProxy(tmp, t)){ return true; }
            if(tmp->left){ sta.emplace(tmp->left); }
            if(tmp->right){ sta.emplace(tmp->right); }
        }
        
        return false;
    }
};

```

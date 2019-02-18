<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7

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
    TreeNode* buildTree(std::vector<int> inorder, std::vector<int> postorder) {
        if(inorder.empty()){
            return nullptr;
        }
        TreeNode* root = new TreeNode(postorder.back());
        auto in_root = std::find(inorder.begin(), inorder.end(), postorder.back());
        auto post_idx = std::next(std::begin(postorder), in_root - std::begin(inorder));
      
        root->left = buildTree(std::vector<int>(inorder.begin(), in_root),
                               std::vector<int>(std::begin(postorder), post_idx));
        root->right = buildTree(std::vector<int>(std::next(in_root), std::end(inorder)),
                                std::vector<int>(post_idx, std::end(postorder) - 1));
        
        return root;
    }
};
```

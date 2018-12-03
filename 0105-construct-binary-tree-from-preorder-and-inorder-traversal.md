<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
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
    TreeNode* buildTree(std::vector<int> preorder, std::vector<int> inorder) {
        if(preorder.empty()){
            return nullptr;
        }
        TreeNode* root = new TreeNode(preorder.front());
        auto in_root = std::find(inorder.begin(), inorder.end(), preorder.front());
        auto pr_idx = std::next(std::begin(preorder),  1 + in_root - std::begin(inorder));
      
        root->left = buildTree(std::vector<int>(std::next(preorder.begin()), pr_idx),
                               std::vector<int>(std::begin(inorder), in_root));
        root->right = buildTree(std::vector<int>(pr_idx, std::end(preorder)),
                                std::vector<int>(std::next(in_root), std::end(inorder)));
        
        return root;
    }
};

```

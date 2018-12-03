<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Example: 

You may serialize the following tree:

    1
   / \
  2   3
     / \
    4   5

as "[1,2,3,null,null,4,5]"
Clarification: The above format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

Note: Do not use class member/global/static variables to store states. Your serialize and deserialize algorithms should be stateless.
```

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Code

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
class Codec {
private:
    int nextVal(std::string str, int& idx){
        auto tmp = idx;
        idx = str.find(' ', idx);
        return std::stoi(str.substr(tmp, idx - tmp + 1));
    }
public:
    //Encodes a tree to a single string.
    std::string serialize(TreeNode* root) {
        if(!root){ return ""; }
        std::string str;
        std::queue<TreeNode*> q;
        q.emplace(root);
        while (!q.empty()) {
            auto p = q.front();
            q.pop();
            if(p){
                q.emplace(p->left);
                q.emplace(p->right);
                str += (std::to_string(p->val) + " ");
            }else{
                str.push_back('#');
            }
        }
        std::cout<<"str:"<<str<<std::endl;
        return str;
    }
    
    // Decodes your encoded data to tree.
    TreeNode* deserialize(std::string data) {
        if(data.empty()){ return nullptr; }
        std::cout<<"data:"<<data<<std::endl;
        
        int idx = 0;
        TreeNode* root = new TreeNode(nextVal(data, idx));
        idx++;
        
        std::queue<TreeNode*> q;
        q.emplace(root);
        
        while(!q.empty()){
            auto p = q.front();
            q.pop();
            if(data[idx] != '#'){
                p->left = new TreeNode(nextVal(data, idx));
                q.emplace(p->left);
            }
            idx++;
            if(data[idx] != '#'){
                p->right = new TreeNode(nextVal(data, idx));
                q.emplace(p->right);
            }
            idx++;
        }
        return root;
    }
};
 
// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));
```

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description
https://leetcode.com/problems/n-ary-tree-preorder-traversal/
Given an n-ary tree, return the preorder traversal of its nodes' values.

For example, given a 3-ary tree:
		   1
		/  |  \
       3   2   4
	  / \
	 5   6
Return its preorder traversal as: [1,3,5,6,2,4].

Note:
Recursive solution is trivial, could you do it iteratively?

- [Recursion](#recursion)
- [Iteration](#iteration)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Recursion

```cpp
/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Solution {
public:
    std::vector<int> preorder(Node* root) {
        if(root == nullptr){ return {}; }
        std::vector<int> result;
        
        std::function<void(Node* root, std::vector<int>& result)>
        preorderProxy =[&](Node* root, std::vector<int>& result){
            result.emplace_back(root->val);
            for(auto& v : root->children){
                preorderProxy(v, result);
            }
        };
        
        preorderProxy(root, result);
        return result;
    }
};

```

# Iteration

```cpp
class Solution {
public:
    std::vector<int> preorder(Node* root) {
        if(root == nullptr){ return {}; }
        std::vector<int> result;
        std::stack<Node*> s;
        s.emplace(root);
        while(!s.empty()){
            auto p = s.top();
            s.pop();
            result.emplace_back(p->val);
            for(auto it = p->children.rbegin(); it != p->children.rend(); it++){
                s.emplace(*it);
            }
        }
        
        return result;
    }
};

```

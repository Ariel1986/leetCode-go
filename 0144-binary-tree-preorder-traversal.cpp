/*
https://leetcode.com/problems/binary-tree-preorder-traversal/

Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

/*
M1:递归
M2:stack
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
private:
    void preorderTraversalProxy(TreeNode* root, std::vector<int>& result) {
        if(root == nullptr){
            return ;
        }
        
        result.emplace_back(root->val);
        preorderTraversalProxy(root->left, result);
        preorderTraversalProxy(root->right, result);
        return ;
    }
    
public:
    std::vector<int> preorderTraversal(TreeNode* root) {
        if(root == nullptr){
            return {};
        }
        
        std::vector<int> result;
        preorderTraversalProxy(root, result);
        
        return result;
    }
};

/////////////////////////////////////////////////////////////////////////////
//左右子树都进栈
class Solution2 {
public:
    std::vector<int> preorderTraversal(TreeNode* root) {
        if(root == nullptr){
            return {};
        }
        
        std::vector<int> result;
        std::stack<TreeNode*> s;
        s.emplace(root);
        TreeNode* pCur = nullptr;
        while(!s.empty()){
            pCur = s.top();
            s.pop();
            result.emplace_back(pCur->val);
            if(pCur->right != nullptr){
                s.emplace(pCur->right);
            }
            if(pCur->left != nullptr){
                s.emplace(pCur->left);
            }
        }
        
        return result;
    }
};
////////////////////////////////////////////////////////////////////////////
//左子树不进栈
class Solution3 {
public:
    std::vector<int> preorderTraversal(TreeNode* root) {
        if(root == nullptr){
            return {};
        }
        
        std::vector<int> result;
        std::stack<TreeNode*> s;
       
        TreeNode* pCur = root;
        while(pCur != nullptr || !s.empty()){
            while(pCur != nullptr){
                result.emplace_back(pCur->val);
                if(pCur->right != nullptr){
                    s.push(pCur->right);
                }
                
                pCur = pCur->left;
            }
            
            if(!s.empty()){
                pCur = s.top();
                s.pop();
            }
        }
        
        return result;
    }
};
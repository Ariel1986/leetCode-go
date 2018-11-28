/*
https://leetcode.com/problems/binary-tree-postorder-traversal/

Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

/*
M1: 递归
M2: stack
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

//M1:递归
class Solution {
private:
    void postorderTraversalProxy(TreeNode* root, std::vector<int>& result) {
        if(root == nullptr){
            return ;
        }
        
        postorderTraversalProxy(root->left, result);
        postorderTraversalProxy(root->right, result);
        result.emplace_back(root->val);
        return ;
    }
    
public:
    std::vector<int> postorderTraversal(TreeNode* root) {
        if(root == nullptr){
            return {};
        }
        
        std::vector<int> result;
        postorderTraversalProxy(root, result);
        
        return result;
    }
};

////////////////////////////////////////////////////////////////////////////
//M2:stack
class Solution {
public:
    std::vector<int> postorderTraversal(TreeNode* root) {
        if(root == nullptr){
            return {};
        }
        
        std::vector<int> result;
        std::stack<std::pair<TreeNode*, bool>> s;
        TreeNode *pCur = root;
        while(pCur != nullptr || !s.empty()){
            if(pCur != nullptr){
                s.push(std::make_pair(pCur, false));
                pCur = pCur->left;
            }else if(!s.empty()){
                auto p = s.top();
                s.pop();
                
                if(p.second){ //右子树访问完成，直接将该节点存入result
                    result.emplace_back(p.first->val);
                }else{ //处理该节点的右子树
                    s.push(std::make_pair(p.first, true));
                    pCur = p.first->right;
                }
            }
        }
        
        return result;
    }
};

////////////////////////////////////////////////////////////////////
//M3: Stack
class Solution3 {
private:
    void postorderTraversalProxy(TreeNode* root, std::stack<TreeNode*>& s){
        while(root != nullptr){
            if(root->right){
                s.emplace(root->right);
            }
            s.emplace(root);
            root = root->left;
        }
    }
public:
    std::vector<int> postorderTraversal(TreeNode* root) {
        if(root == nullptr){
            return {};
        }
        
        std::vector<int> result;
        std::stack<TreeNode*> s;
        
        postorderTraversalProxy(root, s);
        
        while(!s.empty()){
            auto pCur = s.top();
            s.pop();
            
            if(!s.empty() && s.top() == pCur->right){
                auto pRight = s.top();
                s.pop();
                s.emplace(pCur);
                postorderTraversalProxy(pRight, s);
            }else{
                result.emplace_back(pCur->val);
            }
        }
        
        return result;
    }
};

//////////////////////////////////////////////////////////////////////////
//M4: 逆置数组的方式：
class Solution {
public:
    std::vector<int> postorderTraversal(TreeNode* root) {
        if(root == nullptr){
            return {};
        }
        
        std::vector<int> result;
        std::vector<int> vec;
        std::stack<TreeNode*> s;
        s.emplace(root);
        
        while(!s.empty()){
            auto pCur = s.top();
            s.pop();
            vec.emplace_back(pCur->val);
            if(pCur->left != nullptr){
                s.emplace(pCur->left);
            }
            if(pCur->right != nullptr){
                s.emplace(pCur->right);
            }
        }
        
        //逆置数组
        for(int i = vec.size() - 1; i >= 0; i-- ){
            result.emplace_back(vec[i]);
        }
        
        return result;
    }
};

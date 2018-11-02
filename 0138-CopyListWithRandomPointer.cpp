/*
A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.

Return a deep copy of the list.
*/

/**
 * Definition for singly-linked list with a random pointer.
 * struct RandomListNode {
 *     int label;
 *     RandomListNode *next, *random;
 *     RandomListNode(int x) : label(x), next(NULL), random(NULL) {}
 * };
 */

/*
## Linked List
Construct A -> A' -> B -> B' firstly.
*/

//Solution 1:
class Solution {
public:
RandomListNode *copyRandomList(RandomListNode *head) {
    if(head == nullptr){
        return head;
    }
    
    RandomListNode *pCur = head;
    //复制next
    while(pCur != nullptr){
        RandomListNode *pNew = new RandomListNode(pCur->label);
        pNew->next = pCur->next;
        pCur->next = pNew;
        
        pCur = pCur->next->next;
    }
    
    //复制random信息
    pCur = head;
    while(pCur != nullptr){
        if(pCur->random != nullptr){
            pCur->next->random = pCur->random->next;
        }
        pCur = pCur->next->next;
    }
    
    pCur = head;
    while(pCur != nullptr){
        std::cout<<"cur:"<<pCur->label<<std::endl;
        pCur = pCur->next;
    }
    
    //拆分链表
    pCur = head;
    RandomListNode *pResult = head->next;
    RandomListNode *pNew = pResult;
    while(pNew != nullptr && pNew->next != nullptr){
        pCur->next = pCur->next->next;
        pNew->next = pNew->next->next;
        pCur = pCur->next;
        pNew = pNew->next;
    }
    
    //恢复原来链表的最后一个值
    pCur->next = nullptr;
    return pResult;
}
};


// HashTable

class Solution {
public:
RandomListNode* getCloneNode(std::map<RandomListNode*, RandomListNode*>& nodeMap, RandomListNode* head) {
    if (head == nullptr){
        return nullptr;
    }
    
    auto search = nodeMap.find(head);
    if( search == nodeMap.end()){
        nodeMap[head] = new RandomListNode(head->label);
    }
    
    return nodeMap[head];
}

RandomListNode *copyRandomList(RandomListNode *head) {
    if(head == nullptr){
        return nullptr;
    }
    
    std::map<RandomListNode*, RandomListNode*> nodeMap;
    RandomListNode *pNew = new RandomListNode(head->label);
    RandomListNode *pResult = pNew;
    RandomListNode *pOld = head;
    
    while(pOld != nullptr){
        pNew->random = getCloneNode(nodeMap, pOld->random);
        pNew->next = getCloneNode(nodeMap, pOld->next);
        
        pOld = pOld->next;
        pNew = pNew->next;
    }
    
    return pResult;
}
};
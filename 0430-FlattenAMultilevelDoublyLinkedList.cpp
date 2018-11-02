/*

You are given a doubly linked list which in addition to the next and previous pointers, it could have a child pointer, which may or may not point to a separate doubly linked list. These child lists may have one or more children of their own, and so on, to produce a multilevel data structure, as shown in the example below.

Flatten the list so that all the nodes appear in a single-level, doubly linked list. You are given the head of the first level of the list.

Example:

Input:
 1---2---3---4---5---6--NULL
         |
         7---8---9---10--NULL
             |
             11--12--NULL

Output:
1-2-3-7-8-11-12-9-10-4-5-6-NULL
 

Explanation for the above example:

Given the following multilevel doubly linked list:

We should return the following flattened doubly linked list:
详见链接：
https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/
*/

#include <stack>
#include <iostream>
using namespace std;

/*
// Definition for a Node.
*/
class Node {
public:
    int val = NULL;
    Node* prev = NULL;
    Node* next = NULL;
    Node* child = NULL;

    Node() {}

    Node(int _val, Node* _prev, Node* _next, Node* _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
    }
};


//Solution 1:
class Solution {
public:
Node* getTail(Node* head){
    Node *pResp = nullptr;
    while(head != nullptr){
        pResp = head;
        head = head->next;
    }

    return pResp;
}

Node* flatten(Node* head) {
    Node *pCur = head, *pTail = nullptr;

    while(pCur != nullptr){
        if (pCur->child != nullptr){
            Node *originNext = pCur->next;
            pTail = getTail(pCur->child);
            pTail->next = originNext;
            if(originNext != nullptr){
                originNext->prev = pTail;
            }
            
            pCur->next = pCur->child;
            pCur->next->prev = pCur;
            pCur->child = nullptr;
        }

        pCur = pCur->next; 
    }

    return head;
}
};


//Solution 2:
class Solution {
public:
Node* flatten(Node* head) {
        if (head == nullptr){
			return head;
		}

		std::stack<Node*> s;
		Node dummyNode = Node(0, nullptr, head, nullptr);
		Node *pCur = head, *pTail = &dummyNode;
		while ( !s.empty() || pCur != nullptr){
			if (pCur == nullptr) {
				pCur = s.top();
				s.pop();
			}	
			pTail->next = pCur;
			pTail->child = nullptr;
            pCur->prev = pTail;
			pTail = pTail->next;
			if (pCur->child == nullptr){
				pCur = pCur->next;
			}else{
				if (pCur->next != nullptr){
					s.emplace(pCur->next);
				}
				pCur = pCur->child;
			}		
		}

        pTail->child = nullptr;
        pTail->next = nullptr;
		dummyNode.next->prev = nullptr;
		return dummyNode.next;
    }
};


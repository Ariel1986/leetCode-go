package leetCode

/*
https://leetcode.com/problems/linked-list-cycle-ii/

Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

Note: Do not modify the linked list.

Follow up:
Can you solve it without using extra space?
*/

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
 
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
 class Solution {
	public:
			ListNode *detectCycle(ListNode *head) {
				ListNode *pS = head, *pF = head;
				while(pS != nullptr && pF != nullptr && pF->next != nullptr ){
					pS  = pS->next;
					pF = pF->next->next;
					if ( pF == pS){
						break;
					}
				}
				
				if (pF == nullptr || pF->next == nullptr){
					return nullptr;
				}
				
				pF = head;
				while(pF != pS){
					pS = pS->next;
					pF = pF->next;
				}
				
				return pS;
			}
	};
	
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
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	pS, pF := head, head
	for pS != nil && pF != nil && pF.Next != nil {
		pS, pF = pS.Next, pF.Next.Next
		if pS == pF {
			break
		}
	}

	//不存在环, Note： 不可以用pF != pS作为判断条件，对于不存在环的不成立。
	// [1] : 应该返回-1， 如果用pF != pS，会返回1
	if pF == nil || pF.Next == nil {
		return nil
	}

	pF = head
	for pF != pS {
		pF, pS = pF.Next, pS.Next
	}

	return pF
}

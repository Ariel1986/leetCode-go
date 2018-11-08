package leetCode

/*
https://leetcode.com/problems/reverse-nodes-in-k-group/

Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	var (
		resp            = ListNode{Next: head}
		pre, pbeg, pend = &resp, head, head
	)
	for pbeg != nil {
		i := 0
		for i < k && pend != nil {
			pend, i = pend.Next, i+1
		}
		if i < k {
			break
		}
		pTail := pbeg
		for pbeg != pend {
			pbeg.Next, pre.Next, pbeg = pre.Next, pbeg, pbeg.Next
		}
		pTail.Next, pre, pbeg = pend, pTail, pend
	}

	return resp.Next
}

package leetCode

/*
https://leetcode.com/problems/swap-nodes-in-pairs/

Given a linked list, swap every two adjacent nodes and return its head.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
Note:

Your algorithm should use only constant extra space.
You may not modify the values in the list's nodes, only nodes itself may be changed.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		resp = ListNode{Next: head}
		pre  = &resp
	)
	for pre.Next != nil && pre.Next.Next != nil {
		a, b := pre.Next, pre.Next.Next
		pre.Next, a.Next, b.Next = b, b.Next, a
		pre = a
	}
	return resp.Next
}

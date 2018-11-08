package leetCode

/*
https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var (
		result = ListNode{
			Val:  0,
			Next: nil,
		}
		cur   = &result
		carry = 0
	)

	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		carry, val = val/10, val%10

		cur.Next = &ListNode{
			Val: val,
		}
		cur, l1, l2 = cur.Next, l1.Next, l2.Next
	}

	for l1 != nil {
		val := l1.Val + carry
		carry, val = val/10, val%10

		cur.Next = &ListNode{
			Val: val,
		}
		cur, l1 = cur.Next, l1.Next
	}

	for l2 != nil {
		val := l2.Val + carry
		carry, val = val/10, val%10

		cur.Next = &ListNode{
			Val: val,
		}
		cur, l2 = cur.Next, l2.Next
	}

	if carry > 0 {
		cur.Next = &ListNode{
			Val: carry,
		}
	}
	return result.Next
}

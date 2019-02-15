package leetCode

/*
https://leetcode.com/problems/merge-two-sorted-lists/
*/

/*
循环
*/
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	cur := ret
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1, cur = l1.Next, cur.Next
		} else {
			cur.Next = l2
			l2, cur = l2.Next, cur.Next
		}
	}
	for l1 != nil {
		cur.Next = l1
		l1, cur = l1.Next, cur.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2, cur = l2.Next, cur.Next
	}
	return ret.Next
}

/*
递归
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

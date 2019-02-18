package leetCode

/*
https://leetcode.com/problems/sort-list/

Sort a linked list in O(n log n) time using constant space complexity.

Example 1:

Input: 4->2->1->3
Output: 1->2->3->4
Example 2:

Input: -1->5->3->4->0
Output: -1->0->3->4->5
*/

/* Approach
先递归划分子问题，然后合并结果。把待排序列看成由两个有序的子序列，然后合并两个子序列，
然后把子序列看成由两个有序序列。。。。。倒着来看，其实就是先两两合并，然后四四合并。。。
最终形成有序序列。空间复杂度为O(n)，时间复杂度为O(nlogn)。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitList(head *ListNode) *ListNode {
	pSlow, pFast := head, head.Next

	for pFast != nil && pFast.Next != nil {
		pSlow, pFast = pSlow.Next, pFast.Next.Next
	}

	pFast = pSlow.Next
	pSlow.Next = nil

	return pFast
}

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil {
		return l1
	}
	if l1 == nil {
		return l2
	}
	dummyNode := ListNode{}
	cur := &dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummyNode.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := splitList(head)
	return mergeList(sortList(head), sortList(mid))
}

package leetCode

/*
https://leetcode.com/problems/intersection-of-two-linked-lists/
Write a program to find the node at which the intersection of two singly linked lists begins.
Example: 见链接

Notes:
If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* Approach 1:
Brute Force
*/

/* Approach 2:
Hash Table
*/

/* Approach 3:
Two Pointers
Maintain two pointers pApA and pBpB initialized at the head of A and B, respectively. Then let them both traverse through the lists, one node at a time.
When pApA reaches the end of a list, then redirect it to the head of B (yes, B, that's right.); similarly when pBpB reaches the end of a list, redirect it the head of A.
If at any point pApA meets pBpB, then pApA/pBpB is the intersection node.
To see why the above trick would work, consider the following two lists: A = {1,3,5,7,9,11} and B = {2,4,9,11}, which are intersected at node '9'. Since B.length (=4) < A.length (=6), pBpB would reach the end of the merged list first, because pBpB traverses exactly 2 nodes less than pApA does. By redirecting pBpB to head A, and pApA to head B, we now ask pBpB to travel exactly 2 more nodes than pApA would. So in the second iteration, they are guaranteed to reach the intersection node at the same time.
If two lists have intersection, then their last nodes must be the same one. So when pApA/pBpB reaches the end of a list, record the last element of A/B respectively. If the two last elements are not the same one, then the two lists have no intersections.
Complexity Analysis

如果两个链表没有交点，curA和curB会在nil处相遇
Time complexity : O(m+n)O(m+n).

Space complexity : O(1)O(1).
*/

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headA
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headB
		}
	}
	return curA
}

/* Approach 4:
转化成链表是否有环的问题
*/
func getIntersectionNode4(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	for pA.Next != nil {
		pA = pA.Next
	}
	pA.Next = headB

	for pF, pS := headA, headA; pS != nil && pF != nil && pF.Next != nil; {
		pF, pS = pF.Next.Next, pS.Next

		//有环，即有相交点
		if pF == pS {
			pF = headA
			for pF != pS {
				pF, pS = pF.Next, pS.Next
			}
			pA.Next = nil
			return pF
		}
	}

	//没有相交节点，恢复链表 & 返回nil
	pA.Next = nil
	return nil
}

/* Approach 5:
1.求链表长度
2.判断两链表是否相交
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	listLen := func(head *ListNode) int {
		resp := 0
		for ; head != nil; head = head.Next {
			resp++
		}
		return resp
	}

	lA, lB := listLen(headA), listLen(headB)
	if lB > lA {
		headA, headB = headB, headA
		lA, lB = lB, lA
	}

	for i := 0; i < lA-lB; i++ {
		headA = headA.Next
	}

	for headA != nil && headA != headB {
		headA, headB = headA.Next, headB.Next
	}

	return headA
}

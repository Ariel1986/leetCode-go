package leetCode

/*
ou are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

Example:

Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// M1: 逆置链表，然后用2的方法相加，然后再逆置结果链表
// M2: Cheating：在list很长的时候，会有溢出
func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	x1, x2 := uint64(0), uint64(0)
	for l1 != nil {
		x1 = x1*10 + uint64(l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		x2 = x2*10 + uint64(l2.Val)
		l2 = l2.Next
	}

	x := x1 + x2
	result := ListNode{}
	if x == 0 { //[0],[0]这一类情况
		return &result
	}
	//用头插法将每个元素插入链表
	for x != 0 {
		newNode := ListNode{
			Val:  int(x % 10),
			Next: result.Next,
		}
		result.Next = &newNode
		x = x / 10
	}
	return result.Next
}

//M3:Recursion
func AddTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	getListLen := func(l1 *ListNode) int {
		resp := 0
		for l1 != nil {
			resp++
			l1 = l1.Next
		}
		return resp
	}

	addListZero := func(l *ListNode, num int) *ListNode {
		resp := l
		for num > 0 {
			newNode := ListNode{
				Next: resp,
			}
			resp = &newNode
			num--
		}

		return resp
	}

	var combineList func(l1 *ListNode, l2 *ListNode) (int, *ListNode)
	combineList = func(l1 *ListNode, l2 *ListNode) (int, *ListNode) {
		if l1 == nil || l2 == nil {
			return 0, nil
		}
		carry, node := combineList(l1.Next, l2.Next)
		val := l1.Val + l2.Val + carry
		newNode := ListNode{
			Val:  val % 10,
			Next: node,
		}

		return val / 10, &newNode
	}

	len1 := getListLen(l1)
	len2 := getListLen(l2)
	l1 = addListZero(l1, len2-len1)
	l2 = addListZero(l2, len1-len2)

	carry, resp := combineList(l1, l2)
	if carry > 0 {
		newNode := ListNode{
			Val:  carry,
			Next: resp,
		}
		resp = &newNode
	}

	return resp
}

package leetCode

/*
https://leetcode.com/problems/majority-element-ii/

Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Note: The algorithm should run in linear time and in O(1) space.

Example 1:

Input: [3,2,3]
Output: [3]
Example 2:

Input: [1,1,1,3,3,2,2,2]
Output: [1,2]
*/

/* A1:
hash map
*/

/* A2:
sorting
*/

/* A3:
摩尔投票算法
*/
func majorityElement229(nums []int) []int {
	v1, v2, count1, count2 := 0, 0, 0, 0
	for _, v := range nums {
		if v == v1 {
			count1++
		} else {
			if v == v2 {
				count2++
			} else {
				if count1 == 0 {
					v1, count1 = v, 1
				} else {
					if count2 == 0 {
						v2, count2 = v, 1
					} else {
						count1, count2 = count1-1, count2-1
					}
				}
			}
		}
	}

	// 上一轮遍历找出了两个候选人，但是这两个候选人是否均满足票数大于N/3仍然没法确定，需要重新遍历，确定票数
	count1, count2 = 0, 0
	for _, v := range nums {
		if v1 == v {
			count1++
		} else {
			if v2 == v {
				count2++
			}
		}
	}
	ret := make([]int, 0)
	if count1 > len(nums)/3 {
		ret = append(ret, v1)
	}
	if count2 > len(nums)/3 {
		ret = append(ret, v2)
	}
	return ret
}

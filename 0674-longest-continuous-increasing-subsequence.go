package leetCode

/*
https://leetcode.com/problems/longest-continuous-increasing-subsequence/

Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

Example 1:
Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.
Example 2:
Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1.
*/
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	resp, cur, val := 1, 1, nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > val {
			cur++
			val = nums[i]
		} else {
			if cur > resp {
				resp = cur
			}
			cur, val = 1, nums[i]
		}
	}
	if cur > resp { //[1,3,5,7]
		resp = cur
	}
	return resp
}

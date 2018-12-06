<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/maximum-subarray/

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
```

- [Code](#Code)
- [DivideAndConquer](#DivideAndConquer)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Code

```go
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	resp, sum := nums[0], nums[0]
	for i, iEnd := 1, len(nums); i < iEnd; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > resp {
			resp = sum
		}
	}
	return resp
}
```

# DivideAndConquer
```go
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return maxSubArrayProxy(nums, 0, len(nums)-1)
}

func maxSubArrayProxy(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2
	leftMaxSum := maxSubArrayProxy(nums, left, mid)
	rightMaxSum := maxSubArrayProxy(nums, mid+1, right)

	curSum := nums[mid] + nums[mid+1]
	curMaxSum := curSum
	for i := mid - 1; i >= left; i-- {
		curSum += nums[i]
		curMaxSum = max(curMaxSum, curSum)
	}
	curSum = curMaxSum
	for j := mid + 2; j <= right; j++ {
		curSum += nums[j]
		curMaxSum = max(curMaxSum, curSum)
	}

	return max(max(leftMaxSum, rightMaxSum), curMaxSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

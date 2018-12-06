<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/single-number/

Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
```

- [HashTable](#hasnTable)
- [BitManipulation](#bitManipulation)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# BitManipulation

```go
func singleNumber(nums []int) int {
	resp := 0
	for _, v := range nums {
		resp ^= v
	}

	return resp
}
```

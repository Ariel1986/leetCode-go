<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/climbing-stairs/

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

- [Code](#code)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Code

```go
func climbStairs(n int) int {
    if n <= 0 {
		return 0
	}
	d := make([]int, n)
	d[0] = 1
	d[1] = 2
	for i := 2; i < n; i++ {
		d[i] = d[i-1] + d[i-2]
	}

	return d[n-1]
}
```
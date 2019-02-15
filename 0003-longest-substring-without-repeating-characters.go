package leetCode

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

/*
Approach 1: Brute Force
*/
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func allUnique(s string, start, end int) bool {
	m := make(map[byte]int)
	for ; start < end; start++ {
		if _, ok := m[s[start]]; ok {
			return false
		}
		m[s[start]] = 1
	}
	return true
}
func lengthOfLongestSubstring(s string) int {
	ans, n := 0, len(s)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if allUnique(s, i, j) {
				ans = max(ans, j-i)
			}
		}
	}
	return ans
}

/* Approach 2:
Sliding Window
One-pass Hash Table with Two Pointer
*/
func lengthOfLongestSubstring2(s string) int {
	ans, start := 0, 0
	m := make(map[rune]int)
	for i, str := range s {
		//有重复元素
		if idx, ok := m[str]; ok {
			start = max(idx+1, start)
		}
		ans = max(ans, i-start+1)
		m[str] = i

	}

	return ans
}

/* Approach 3:
Assuming ASCII 128
*/
func lengthOfLongestSubstring3(s string) int {
	ans := 0
	idx := make([]int, 128)
	start := 0
	for i, v := range s {
		start = max(start, idx[v])
		ans = max(ans, i-start+1)
		idx[v] = i + 1
	}

	return ans
}

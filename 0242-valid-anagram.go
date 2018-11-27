package leetCode

/*
https://leetcode.com/problems/valid-anagram/

Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?

*/

/* Method:
1. 排序两个字符串，然后比较字符串是否相等
2. 对字符串建立hash，然后比较hash表是否相等
3. 对字符串s建立hash，遍历字符串t的时候去减hash的值；对于字母的情况，也可以用[26]大小的数组
*/

//用数组实现
func isAnagram(s string, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return true
	}
	if len(s) != len(t) {
		return false
	}

	arr := make([]int, 26)
	for _, v := range s {
		arr[v-'a']++
	}

	for _, v := range t {
		if arr[v-'a'] == 0 {
			return false
		} else {
			arr[v-'a']--
		}
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}

//用map，速度比较慢
func isAnagram2(s string, t string) bool {
	m := make(map[rune]int, 0)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for _, v := range t {
		if _, ok := m[v]; ok {
			m[v]--
			if m[v] == 0 {
				delete(m, v)
			}
		} else {
			return false
		}
	}

	return len(m) == 0
}

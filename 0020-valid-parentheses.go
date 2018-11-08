package leetCode

/*
https://leetcode.com/problems/valid-parentheses/

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	var (
		mem   = map[rune]rune{')': '(', '}': '{', ']': '['}
		stack = make([]rune, 0)
	)

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v) //push
		case ')', '}', ']':
			if len(stack) <= 0 || stack[len(stack)-1] != mem[v] {
				return false
			} else {
				stack = stack[:len(stack)-1] //pop
			}
		default: //其他字符，返回false
			return false
		}
	}

	return len(stack) == 0
}

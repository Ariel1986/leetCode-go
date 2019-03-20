package leetCode

import "fmt"

/*
https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

/*
 */
/* Test case
["flower","flow","flight"]
["aa","a"]
*/
//Approach 1: Horizontal scanning
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		if strs[i] == res {
			continue
		}
		j := 0
		for j < len(res) && j < len(strs[i]) && res[j] == strs[i][j] {
			j++
		}
		res = res[:j]
	}
	return res
}

//Approach 2: Vertical scanning
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[j][:i]
			}
		}
	}

	return strs[0]
}

// Approach 3: Divide and conquer
func commonPrefix(s, r string) string {
	min := len(s)
	if min > len(r) {
		min = len(r)
	}

	for i := 0; i < min; i++ {
		if r[i] != s[i] {
			return s[:i]
		}
	}
	return s[:min]
}

func longestCommonPrefixHelper(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	} else {
		mid := (l + r) / 2
		left := longestCommonPrefixHelper(strs, l, mid)
		right := longestCommonPrefixHelper(strs, mid+1, r)
		return commonPrefix(left, right)
	}
}

func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return longestCommonPrefixHelper(strs, 0, len(strs)-1)
}

//Approach 4: Binary search
func isCommonPrefix(strs []string, length int) bool {
	str := strs[0][:length]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(str); j++ {
			if strs[i][j] != str[j] {
				return false
			}
		}
	}
	return true
}

func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	min_len := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < min_len {
			min_len = len(strs[i])
		}
	}

	low, high := 1, min_len

	for low <= high {
		mid := (low + high) / 2
		if isCommonPrefix(strs, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return strs[0][:(low+high)/2]
}

//Approach 5: Trie
//Trie 参考0208
func longestCommonPrefix5(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	root := Constructor()
	for _, str := range strs {
		root.Insert(str)
	}
	return root.SearchLongestPrefix(strs[0])
}

const MAX_LETTERS = 26

type TrieNode struct {
	isEnd bool
	next  *[MAX_LETTERS]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		next: new([MAX_LETTERS]*TrieNode),
	}
}

////////////////////////////////////////////////////////////
type Trie struct {
	node *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		node: NewTrieNode(),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this.node
	for _, w := range word {
		idx := w - 'a'
		if p.next[idx] == nil {
			p.next[idx] = NewTrieNode()
		}
		p = p.next[idx]
	}
	p.isEnd = true
}
func (this *Trie) SearchLongestPrefix(str string) string {
	res := ""
	p := this.node
	fmt.Println("str: ", str)
	for _, s := range str {
		idx := int(s - 'a')
		if p.next[idx] != nil && !p.isEnd {
			for i := 0; i < MAX_LETTERS; i++ {
				if i == idx {
					continue
				}
				if p.next[i] != nil {
					return res
				}
			}
			res = res + string(s)
			p = p.next[idx]
		} else {
			return res
		}
	}
	return res
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix5(strs)
	fmt.Println(res)
}

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description

```
https://leetcode.com/problems/longest-consecutive-sequence/

Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

Example:

Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
```

- [Approach 1: Brute Force]
```
Time complexity : O(n^3)
```
- [Approach 2: Sorting]
```
Time complexity : O(nlgn).
```
- [Approach 3: HashSet and Intelligent Sequence Building](#approach3)
```
Time complexity : O(n).
```
<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Approach3

```cpp
class Solution {
public:
    int longestConsecutive(std::vector<int>& nums) {
        std::set<int> s;
        for (auto num : nums){
            s.emplace(num);
        }
        
        int longestStreak = 0;
        for(auto num : nums){
            if (s.find(num-1) == s.end()){
                int curNum = num, curStreak = 1;
                while(s.find(curNum+1)!= s.end()){
                    curStreak++;
                    curNum++;
                }
                
                longestStreak = std::max(longestStreak, curStreak);
            }
        }
        
        return longestStreak;
    }
};
```

```go
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	resp := 0
	for k, _ := range m {
		cur := 1
		for {
			_, ok := m[k+1]
			if ok {
				cur++
				k++
			} else {
				break
			}
		}
		if cur > resp {
			resp = cur
		}
	}

	return resp
}
```
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

# Description
https://leetcode.com/problems/kth-largest-element-in-an-array/

Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note: 
You may assume k is always valid, 1 ≤ k ≤ array's length.
```
```

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

```cpp
class Solution {
public:
    int findKthLargest(std::vector<int>& nums, int k) {
        std::multiset<int> s;
        for(auto num: nums){
            s.emplace(num);
            if(s.size() > k){
                s.erase(s.begin());
            }
        }
        
        return *s.begin();
    }
};
```

```go
func findKthLargest(nums []int, k int) int {
	if len(nums) < k && k < 1 {
		return -1
	}

	arr := nums[0:k]
	sort.Ints(arr)
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > arr[k-1] {
			j := k - 1
			for ; j > 0 && arr[j-1] < nums[i]; j-- {
				arr[j] = arr[j-1]
			}
			arr[j] = nums[i]
		}
	}

	return arr[k-1]
}
```
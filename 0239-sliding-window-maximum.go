package leetCode

/*
https://leetcode.com/problems/sliding-window-maximum/

Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
*/

//Method1: deque (double-ended queue)
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return nil
	}

	var (
		resp   []int
		window []int
	)

	for i, v := range nums {
		//1. 如果window的大小超过k，删除窗口的第一个元素
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}

		//2. 删除window中比v老且小的元素
		var tmp []int
		for j, jEnd := 0, len(window); j < jEnd; j++ {
			if nums[window[j]] > v {
				tmp = append(tmp, window[j])
			}
		}
		window = tmp

		//3. 将窗口向后移，即将当前元素放入窗口
		window = append(window, i)

		//4. 将窗口的第一个元素，即当前窗口中最大的元素存入结果
		if i >= k-1 {
			resp = append(resp, nums[window[0]])
		}
	}

	return resp
}

//Method2: Heap
class Solution {
	public:
		std::vector<int> maxSlidingWindow(std::vector<int>& nums, int k) {
			std::vector<int> res;
			if(nums.empty() || k <= 0){
				return res;
			}
			auto cmp = [&](int i, int j){
				return nums[i] < nums[j];
			};
			
			std::vector<int> val;
			
			for(int i = 0, iEnd = nums.size(); i < iEnd; i++){
				val.emplace_back(i);
				
				//heap_pop的方式和push的比较方法不一样
				if(val.size() > k){
					std::make_heap(val.begin(), val.end(),std::greater<int>() );
					std::pop_heap(val.begin(), val.end(),std::greater<int>());
					val.pop_back();
				}
			
				std::make_heap(val.begin(), val.end(), cmp);
				if(val.size() == k){
					res.emplace_back(nums[val.front()]);
				}
			}
			
			return res;
		}
	};
			
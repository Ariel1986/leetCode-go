package leetCode

import "sort"

/*
https://leetcode.com/problems/3sum/

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	ret := make([][]int, 0)
	for i, v := range nums[:len(nums)-2] {
		if i > 0 && v == nums[i-1] { //省略重复元素
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+v > 0 {
				r--
			} else if nums[l]+nums[r]+v < 0 {
				l++
			} else {
				ret = append(ret, []int{v, nums[l], nums[r]})
				//去掉重复元素
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				//移动l和r到适当位置
				l++
				r--
			}
		}
	}
	return ret
}

/* cpp
class Solution {
public:
    std::vector<std::vector<int>> threeSum(std::vector<int>& nums) {
        if(nums.size() < 3){
            return std::vector<std::vector<int>>({});
        }

        //避免重复元素
        std::sort(nums.begin(), nums.end());


        std::vector<std::vector<int>> res;
        for(int i = 0, iEnd = nums.size()-2; i < iEnd; i++){
            if(i > 0 && nums[i] == nums[i-1]){
                continue;
            }

            for(int j = i+1, jEnd = nums.size()-1; j < jEnd; j++){
                if(j > i+1 && nums[j] == nums[j-1]){
                    continue;
                }

                std::set<int> s;
                for(int k = j+1, kEnd = nums.size(); k < kEnd; k++){
                    s.insert(nums[k]);
                }

                if(s.find(0-nums[i]-nums[j]) != s.end()){
                    res.emplace_back(std::vector<int>{nums[i],nums[j],0-nums[i]-nums[j]});
                }
            }
        }

        return res;
    }
};
*/

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

/* Method:
1. bruce force
2. two pointer
3. set
*/

//M2: two pointer
class Solution {
	public:
		std::vector<std::vector<int>> threeSum(std::vector<int>& nums) {
			if(nums.size() < 3){
				return std::vector<std::vector<int>>({});
			}
			
			std::vector<std::vector<int>> res;
			std::sort(nums.begin(), nums.end());
			
			for(int i = 0, iEnd = nums.size() - 2; i < iEnd; i++){
				if(i > 0 && nums[i] == nums[i-1]){
					continue;
				}
				int l = i+1, r = nums.size()-1;
				while(l < r){
					int sum = nums[i] + nums[l] + nums[r];
					if(sum < 0){
						l++;
					}else{
						if(sum > 0){
							r--;
						}
						else{
							res.emplace_back(std::vector<int>{nums[i], nums[l], nums[r]});
							while( l < r && nums[l] == nums[l+1]){
								l++;
							}
							while(l<r && nums[r] == nums[r-1]){
								r--;
							}
							l++;
							r--;
						}
					}
				}
			}
			
			return res;
		}
	};


//M3: set
class Solution {
	public:
		std::vector<std::vector<int>> threeSum(std::vector<int>& nums) {
			if(nums.size() < 3){
				return std::vector<std::vector<int>>({});
			}
			
			//避免重复元素
			//std::sort(nums.begin(), nums.end());
			
			std::set<std::vector<int>> res;
			for(int i = 0, iEnd = nums.size()-2; i < iEnd; i++){
				if(i > 0 && nums[i] == nums[i-1]){
					continue;
				}
				
				std::set<int> s;
				for(int j = i+1, jEnd = nums.size(); j < jEnd; j++){
					if(s.find(0-nums[i]-nums[j]) != s.end()){
						res.insert(std::vector<int>{nums[i],0-nums[i]-nums[j], nums[j]});
					}else{
						s.insert(nums[j]);
					}
				}
				
			}
			
			std::vector<std::vector<int>> vec;
			
			for(auto&& v: res){
				vec.emplace_back(v);
			}
			return vec;
		}
	};
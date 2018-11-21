/*
https://leetcode.com/problems/sliding-window-median/

Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

Examples: 
[2,3,4] , the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Your job is to output the median array for each window in the original array.

For example,
Given nums = [1,3,-1,-3,5,3,6,7], and k = 3.

Window position                Median
---------------               -----
[1  3  -1] -3  5  3  6  7       1
 1 [3  -1  -3] 5  3  6  7       -1
 1  3 [-1  -3  5] 3  6  7       -1
 1  3  -1 [-3  5  3] 6  7       3
 1  3  -1  -3 [5  3  6] 7       5
 1  3  -1  -3  5 [3  6  7]      6
Therefore, return the median sliding window as [1,-1,-1,3,5,6].

Note: 
You may assume k is always valid, ie: k is always smaller than input array's size for non-empty array.
*/

class Solution {
public:
    std::vector<double> medianSlidingWindow(std::vector<int>& nums, int k) {
        std::vector<double> res;
        if(nums.empty() || k <= 0){
            return res;
        }
        
        std::multiset<int> l, r;
        
        for(int i = 0, iEnd = nums.size(); i < iEnd; i++){
            //1. insert new member
            if(r.empty() || nums[i] > *r.begin()){
                r.emplace(nums[i]);
            }else{
                l.emplace(nums[i]);
            }
            
            //2. Remove number out of the window
            if(i >= k){
                if(nums[i-k] >= *r.begin()){
                    r.erase(r.find(nums[i-k]));
                }else{
                    l.erase(l.find(nums[i-k]));
                }
            }
            
            //3. Balance the size of two sets
            while(r.size()  < l.size()){
                r.insert(*l.rbegin());
                l.erase(--l.end());
            }
            
            while( r.size() > l.size() + 1){
                l.insert(*r.begin());
                r.erase(r.begin());
            }
            
            
            //4. push back median
            if(i >= k-1){
                if(k&1) {
                    res.emplace_back(*r.begin());
                }else{
                    res.emplace_back((double(*r.begin()) + *l.rbegin())/2);
                }
            }
        }

        return res;
    }
};
        
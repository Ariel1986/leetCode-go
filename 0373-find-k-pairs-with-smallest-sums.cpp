package leetcode

/*
https://leetcode.com/problems/find-k-pairs-with-smallest-sums/

You are given two integer arrays nums1 and nums2 sorted in ascending order and an integer k.

Define a pair (u,v) which consists of one element from the first array and one element from the second array.

Find the k pairs (u1,v1),(u2,v2) ...(uk,vk) with the smallest sums.

Example 1:

Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
Output: [[1,2],[1,4],[1,6]] 
Explanation: The first 3 pairs are returned from the sequence: 
             [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
Example 2:

Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
Output: [1,1],[1,1]
Explanation: The first 2 pairs are returned from the sequence: 
             [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
Example 3:

Input: nums1 = [1,2], nums2 = [3], k = 3
Output: [1,3],[2,3]
Explanation: All possible pairs are returned from the sequence: [1,3],[2,3]
*/

class Solution {
public:
    std::vector<std::pair<int, int>> kSmallestPairs(std::vector<int>& nums1, std::vector<int>& nums2, int k) {
        std::vector<std::pair<int, int>> result;
        if(nums1.empty() || nums2.empty() || k <=0){
            return result;
        }
        
        auto cmp = [](std::tuple<int, int, int>& n1, std::tuple<int, int, int>& n2){
            return std::get<0>(n1) > std::get<0>(n2);
        };
        
        std::priority_queue<std::tuple<int, int, int>,std::vector<std::tuple<int, int, int>> , decltype(cmp)> pq(cmp);

        for(int i = 0, iEnd = nums2.size(); i < iEnd; i++){
            pq.emplace(std::make_tuple(nums1[0]+nums2[i], 0, i));
        }

        for(int i = 0, iEnd = std::min(k, int(nums1.size() * nums2.size())); i < iEnd; i++){
            std::tuple<int, int, int> val = pq.top();
            pq.pop();
            int i1 =std::get<1>(val), i2=std::get<2>(val);
            std::cout<<std::get<0>(val)<<","<<i1<<","<<i2<<std::endl;
            result.emplace_back(std::make_pair(nums1[i1], nums2[i2]));
            if(i1 == nums1.size()-1){
                continue;
            }
            pq.emplace(std::make_tuple(nums1[i1+1] + nums2[i2], i1+1, i2));
        }
        
        return result;
    }
};


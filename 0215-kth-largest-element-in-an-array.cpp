package leetCode

/*
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
*/

//M1: use heap
class Solution {
	public:
		int findKthLargest(std::vector<int>& nums, int k) {
			std::vector<int> vec;
			for(int i = 0, iEnd = std::min(int(nums.size()), k); i < iEnd; i++){
				vec.emplace_back(nums[i]);
			}
			std::make_heap(vec.begin(), vec.end(), std::greater<int>());
			for(int i = k, iEnd = nums.size(); i < iEnd; i++){
				vec.emplace_back(nums[i]);
				std::push_heap(vec.begin(), vec.end(), std::greater<int>());
				std::pop_heap(vec.begin(),vec.end(), std::greater<int>());
				vec.pop_back();
			}
			
			return vec[0];
		}
	};

//M2: https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/60309/C%2B%2B-PartitionMax-Heappriority_queuemultiset
//Partation:
//This problem needs to use partial sorting algorithms. There is a built-in nth_element for this.
class Solution {
public:
int findKthLargest(std::vector<int>& nums, int k) {
	std::nth_element(nums.begin(), nums.begin()+k-1, nums.end(), std::greater<int>());
	return nums[k-1];
}
};

/*
Then let's implement the partial sorting algorithms by ourselves, using quicksort (partition) and heapsort.

Quicksort
In quicksort, in each iteration, we need to select a pivot and then partition the array into roughly two halves:

Elements larger than or equal to the pivot;
Elements smaller than or equal to the pivot.
First let's do an example with the array [3, 2, 1, 5, 4, 6] in the problem statement. Let's assume in each time we select the leftmost element to be the pivot, in this case, 3. We then use it to partition the array, which results in [5, 6, 4, 3, 1, 2]. Now 3 is in the 3rd (0-based) position and we know that it is the 4th (1-based) largest element. So, have you noticed how partition is related to this problem?

In the above example, now we know that 3 is the 4th largest element. If we are asked to find the 2nd largest element, it should be to the left of 3. If we are asked to find the 5th largest element, it should be to the right of 3. So, in the average sense, the problem is reduced to approximately half of its original size, giving the recursion T(n) = T(n/2) + O(n) in which O(n) is the time for partition. This recursion, once solved, gives T(n) = O(n) and thus we have a linear time solution. Note that since we only need to consider one half of the array, the time complexity is O(n). If we need to consider both the two halves of the array, like quicksort, then the recursion will be T(n) = 2T(n/2) + O(n) and the complexity will be O(nlogn).

Of course, O(n) is the average/best time complexity. In the worst case, the recursion may become T(n) = T(n - 1) + O(n) and the complexity will be O(n^2).

The algorithm is as follows.

Initialize left to be 0 and right to be nums.size() - 1;
Partition the array, if the pivot is at the k-1-th position, return it (we are done);
If the pivot is right to the k-1-th position, update right to be the left neighbor of the pivot;
Else update left to be the right neighbor of the pivot.
Repeat 2.
And the code is as follows.
*/
class Solution {
	private:
		int partition(std::vector<int>& nums, int left, int right){
			int l = left+1, r = right, pivot = nums[left];
			while(l <= r){
				if(nums[l] < pivot && nums[r] >pivot){
					std::swap(nums[l++], nums[r--]);
				}
				if(nums[l] >= pivot){
					l++;
				}
				if(nums[r] <= pivot){
					r--;
				}
			}
			std::swap(nums[left], nums[r]);
			return r;
		}
		
	public:
		int findKthLargest(std::vector<int>& nums, int k) {
			int left = 0, right = nums.size()-1;
			while(true){
				int p = partition(nums, left, right);
				if (p == k-1){
					return nums[p];
				}
				if(p > k-1){
					right = p-1;
				}else{
					left = p+1;
				}
			}
		}
	};

//M3: priority_deque
//For heapsort, we may also use the built-in priority_queue. This is a max-heap by default.
//占用时间比较少
class Solution {
	public:
		int findKthLargest(std::vector<int>& nums, int k) {
			std::priority_queue<int> pq(nums.begin(), nums.end());
			for(int i = 0; i < k-1; i++){
				pq.pop();
			}
			return  pq.top();
		}
};

//We may also use priority_queue as a min-heap and the code will become.
class Solution {
	public:
		int findKthLargest(std::vector<int>& nums, int k) {
			std::priority_queue<int, std::vector<int>, std::greater<int>> pq;
			for(auto num : nums){
				pq.emplace(num);
				if(pq.size() > k){
					pq.pop();
				}
			}
			return  pq.top();
		}
};

//For the min-heap, we can also use multiset.
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
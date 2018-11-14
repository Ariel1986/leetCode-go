package leetCode

/*
https://leetcode.com/problems/kth-largest-element-in-a-stream/
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

Example:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
Note:
You may assume that nums' length ≥ k-1 and k ≥ 1.
*/

// M1: use heap
class KthLargest {
	private:
		std::vector<int> vec;
		int k;
	public:
		KthLargest(int k, std::vector<int> nums) {
			for(int i = 0, iEnd = nums.size(); i < k && i < iEnd; i ++){
				vec.emplace_back(nums[i]);
			}
			
			//元素由小->大排序
			std::make_heap(vec.begin(), vec.end(),[](int i, int j){return i > j;});
		   
			for(int i = k, iEnd = nums.size(); i < iEnd; i++){
				if(nums[i] > vec.front()){
					vec.front() = nums[i];
					std::make_heap(vec.begin(), vec.end(),[](int i, int j){return i > j;});
				}
			}
			this->k = k;
		}
		
		int add(int val) {
			if(vec.size() < k){
				vec.emplace_back(val);
				std::make_heap(vec.begin(), vec.end(),[](int i, int j){return i > j;});
			}else{
				if(val > vec.front()){
					vec.front() = val;
					std::make_heap(vec.begin(), vec.end(),[](int i, int j){return i > j;});
				}
			}
			
			return vec.front();
		}
	};

// M2: use priority_queue
class KthLargest {
	private:
		std::priority_queue<int,std::vector<int>, std::greater<int>> dq;
		int K;
	public:
		KthLargest(int k, std::vector<int> nums) {
			K = k;
			for(auto n : nums){
				add(n);
			}
		}
		
		int add(int val) {
			if(dq.size() < K || val > dq.top()){
				dq.push(val);
			}
			
			if(dq.size() > K){
				dq.pop();
			}
			
			return dq.top();
		}
	};
	/**
	 * Your KthLargest object will be instantiated and called as such:
	 * KthLargest obj = new KthLargest(k, nums);
	 * int param_1 = obj.add(val);
	 */

// M3: Go method
//  https://leetcode.com/problems/kth-largest-element-in-a-stream/discuss/179019/Min-Heap-implementation-in-Go

package leetCode

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func (h *IntHeap) Pop() interface{} {
	result := (*h)[(*h).Len()-1]
	*h = (*h)[:(*h).Len()-1]
	return result
}

type KthLargest struct {
	Nums *IntHeap
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	result := KthLargest{
		K:    k,
		Nums: &IntHeap{},
	}
	heap.Init(result.Nums)
	for _, v := range nums {
		result.Add(v)
	}
	return result
}

func (this *KthLargest) Add(val int) int {
	if len(*this.Nums) < this.K {
		heap.Push(this.Nums, val)
	} else if val > (*this.Nums)[0] {
		heap.Push(this.Nums, val)
		heap.Pop(this.Nums)
	}
	return (*this.Nums)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

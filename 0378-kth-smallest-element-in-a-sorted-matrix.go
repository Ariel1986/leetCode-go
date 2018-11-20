package leetCode

/*
   https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
   Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.

   Note that it is the kth smallest element in the sorted order, not the kth distinct element.

   Example:

   matrix = [
      [ 1,  5,  9],
      [10, 11, 13],
      [12, 13, 15]
   ],
   k = 8,

   return 13.
*/

//https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/discuss/85173/Share-my-thoughts-and-Clean-Java-Code

/*
   Solution 1 : Heap
   Here is the step of my solution:

   Build a minHeap of elements from the first row.
   Do the following operations k-1 times :
   Every time when you poll out the root(Top Element in Heap), you need to know the row number and column number of that element(so we can create a tuple class here), replace that root with the next element from the same column.
   After you finish this problem, thinks more :

   For this question, you can also build a min Heap from the first column, and do the similar operations as above.(Replace the root with the next element from the same row)
   What is more, this problem is exact the same with Leetcode373 Find K Pairs with Smallest Sums, I use the same code which beats 96.42%, after you solve this problem, you can check with this link:
   https://discuss.leetcode.com/topic/52953/share-my-solution-which-beat-96-42
*/
class Solution {
	public:
		int kthSmallest(std::vector<std::vector<int>>& matrix, int k) {
			std::vector<std::tuple<int, int, int>> vec;
			for(int i = 0, iEnd = matrix.size(); i < iEnd; i++){
				vec.emplace_back(std::make_tuple(matrix[i][0], i, 0));
			}
			
			auto cmp = [](std::tuple<int, int, int> a, std::tuple<int,int, int> b){
				return std::get<0>(a) > std::get<0>(b);
			};
			
			std::make_heap(vec.begin(), vec.end(), cmp);
			
			for(int i = 1; i < k; i++){
				std::pop_heap(vec.begin(), vec.end(), cmp);
				auto val = vec.back();
				
				std::cout << std::get<0>(val) <<"," <<std::get<1>(val) << "," << std::get<2>(val)<<std::endl;
				vec.pop_back();
				if(std::get<2>(val)+ 1 < matrix[std::get<1>(val)].size()){
					vec.emplace_back(std::make_tuple(matrix[std::get<1>(val)][std::get<2>(val)+ 1], std::get<1>(val),std::get<2>(val)+1));
					std::push_heap(vec.begin(), vec.end(), cmp);
				}
			}
			
			return  std::get<0>(vec[0]);
		}
};
/*
   Solution 2 : Binary Search
   We are done here, but let's think about this problem in another way:
   The key point for any binary search is to figure out the "Search Space". For me, I think there are two kind of "Search Space" -- index and range(the range from the smallest number to the biggest number). Most usually, when the array is sorted in one direction, we can use index as "search space", when the array is unsorted and we are going to find a specific number, we can use "range".

   Let me give you two examples of these two "search space"

   index -- A bunch of examples -- https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/ ( the array is sorted)
   range -- https://leetcode.com/problems/find-the-duplicate-number/ (Unsorted Array)
   The reason why we did not use index as "search space" for this problem is the matrix is sorted in two directions, we can not find a linear way to map the number and its index.
*/

func kthSmallest(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	low, high := matrix[0][0], matrix[m-1][n-1]+1

	for low <= high {
		//mid := low + (high-low)/2
		mid := (low + high) / 2
		count, j := 0, n-1

		for i := 0; i < m; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			count += (j + 1)
		}

		if count < k {
			low = mid + 1
		} else {
			//high = mid
			high = mid - 1
		}
	}

	return low
}

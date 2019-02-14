package leetCode

import (
	"math/rand"
	"sort"
)

/*
https://leetcode.com/problems/majority-element/

Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2
*/

/* Method:
https://leetcode.com/problems/majority-element/discuss/51612/6-Suggested-Solutions-in-C%2B%2B-with-Explanations
*/

/* Method 1
Brute Force
*/

/* Method 2
Hash Table
*/

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}

	return -1
}

/* method 3:
Sorting
*/
func majorityElement3(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return nums[len(nums)/2]
}

/* Method 4:
Randomization 随机化
随机的取一个数，看它是不是众数
*/
func majorityElement4(nums []int) int {
	length := int32(len(nums))
	for {
		idx := rand.Int31n(length)
		count := 0
		for _, v := range nums {
			if v == nums[idx] {
				count++
			}
			if count > len(nums)/2 {
				return nums[idx]
			}
		}
	}
	return -1
}

/* Method 5:
Divide and Conquer
*/
func countInRange(nums []int, num, low, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if nums[i] == num {
			count++
		}
	}
	return count
}

func majorityElementRec(nums []int, low, high int) int {
	// nums只有一个元素
	if low == high {
		return nums[low]
	}

	mid := low + (high-low)/2
	left := majorityElementRec(nums, low, mid)
	right := majorityElementRec(nums, mid+1, high)

	//nums左右两部分有相同的众数，就返回该结果。
	if left == right {
		return left
	}

	// otherwise, 分别计算数组左右两边的众数的量，并返回量最多的数
	leftCount := countInRange(nums, left, low, mid)
	rightCount := countInRange(nums, right, mid+1, high)
	if leftCount > rightCount {
		return left
	}
	return right
}

func majorityElement5(nums []int) int {
	return majorityElementRec(nums, 0, len(nums)-1)
}

/* Method 6:
Boyer-Moore Voting Algorithm 摩尔投票算法
https://blog.csdn.net/u014248127/article/details/79230221
算法原理：
每次从数组中找出一对不同的元素，将他们从数组中删除，直到遍历完整个数组。由于题目明确指出一定存在一个出现次数超过一半的元素，
所以遍历完数组后一数组中一定存在至少一个元素。
1. 定义一个序列原序m和一个计数器count，并初始化count=0
2. 依次扫描数组中的元素v，如果count=0，令m=v
3. 如果count不等于0，比较m和v，如果m==v，则count+1； 否则count-1
4. 遍历完后，m的值就是众数
*/

func majorityElement6(nums []int) int {
	m, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			m = v
			count++
		} else {
			if m == v {
				count++
			} else {
				count--
			}
		}
	}

	return m
}

/* Method 7:
Bit Manipulation
统计nums中每个数每一位的1的个数，如果该位为1的量>len(nums)/2, 结果的该位为1.
对于有负数的情况该方法不可用
*/
func majorityElement7(nums []int) int {
	result, mask, n := 0, 1, len(nums)/2
	for i := 0; i < 32; i++ {
		bitCount := 0
		for _, v := range nums {
			if v&mask != 0 {
				bitCount++
			}
			if bitCount > n {
				result |= mask
				break
			}
		}
		mask <<= 1
	}

	return result
}

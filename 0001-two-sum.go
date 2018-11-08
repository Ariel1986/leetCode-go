package leetCode

import "sort"

/*
https://leetcode.com/problems/two-sum/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

//M1: Brute Force
func twoSum(nums []int, target int) []int {
	nLen := len(nums)
	for i, v := range nums {
		for j := i + 1; j < nLen; j++ {
			if nums[j] == target-v {
				return []int{i, j}
			}
		}
	}
	return nil
}

//M2: Sort and Two Pointers
func twoSum2(nums []int, target int) []int {
	type pair struct {
		val   int
		index int
	}
	pairs := make([]pair, 0)
	for i, v := range nums {
		pairs = append(pairs, pair{
			val:   v,
			index: i,
		})
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].val < pairs[j].val })

	for beg, end := 0, len(pairs)-1; beg < end; {
		if pairs[beg].val+pairs[end].val < target {
			beg++
		} else if pairs[beg].val+pairs[end].val > target {
			end--
		} else {
			return []int{pairs[beg].index, pairs[end].index}
		}
	}

	return nil
}

//M3: Two-pass Hash Table
// Trick Points:
//Consider some edge cases, such [3, 3], 6. The index of answers should be different.
//[3, 2, 4] 6
// [1, 2]
func twoSum3(nums []int, target int) []int {
	var (
		hash = make(map[int]int)
	)
	for i, v := range nums {
		hash[v] = i
	}

	for i, v := range nums {
		if j, ok := hash[target-v]; ok && i != j {
			return []int{i, hash[target-v]}
		}
	}

	return nil
}

//M4: One-pass Hash Table
// Following the last solution, two-pass is not necessary, because when you visit the current item, you just need to query the elements before it.
func twoSum4(nums []int, target int) []int {
	var (
		hash = make(map[int]int)
	)
	for i, v := range nums {
		if _, ok := hash[target-v]; ok {
			if i < hash[target-v] {
				return []int{i, hash[target-v]}
			} else {
				return []int{hash[target-v], i}
			}
		}
		hash[v] = i
	}

	return nil
}

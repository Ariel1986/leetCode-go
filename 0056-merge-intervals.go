package leetCode

import "sort"

/*
https://leetcode.com/problems/merge-intervals/
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })

	resp := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if rEnd := resp[len(resp)-1]; rEnd.End >= intervals[i].Start {
			if intervals[i].End > rEnd.End {
				resp[len(resp)-1].End = intervals[i].End
			}
		} else {
			resp = append(resp, intervals[i])
		}
	}

	return resp
}

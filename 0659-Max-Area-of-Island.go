package leetCode

import "reflect"

/*
https://leetcode.com/problems/max-area-of-island/

Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

Example 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
Example 2:

[[0,0,0,0,0,0,0,0]]
Given the above grid, return 0.
Note: The length of each dimension in the given grid does not exceed 50.
*/

/*
Approach #1: Depth-First Search (Recursive)
*/

func area(grid [][]int, seen [][]bool, r, c int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || seen[r][c] || grid[r][c] == 0 {
		return 0
	}
	seen[r][c] = true

	return 1 + area(grid, seen, r+1, c) + area(grid, seen, r-1, c) + area(grid, seen, r, c+1) + area(grid, seen, r, c-1)
}

func maxAreaOfIsland(grid [][]int) int {
	seen := make([][]bool, len(grid))
	for i := range seen {
		seen[i] = make([]bool, len(grid[0]))
	}

	ans := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			resp := area(grid, seen, r, c)
			if resp > ans {
				ans = resp
			}
		}
	}

	return ans
}

/*
Approach #2: Depth-First Search (Iterative)
*/

func maxAreaOfIsland2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	seen := make([][]bool, len(grid))
	for i := range seen {
		seen[i] = make([]bool, len(grid[0]))
	}

	ans := 0
	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 && !seen[r][c] {
				cur := 0
				seen[r][c] = true
				s := New() //9000 stack.go
				s.Push([]int{r, c})
				for s.Len() > 0 {
					node := s.Pop()
					n := reflect.ValueOf(node)
					nr, nc := int(n.Index(0).Int()), int(n.Index(1).Int())
					cur++
					for k := 0; k < 4; k++ {
						cr, cc := nr+dr[k], nc+dc[k]
						if 0 <= cr && cr < len(grid) &&
							0 <= cc && cc < len(grid[0]) &&
							grid[cr][cc] == 1 && !seen[cr][cc] {
							s.Push([]int{cr, cc})
							seen[cr][cc] = true
						}
					}
				}
				if cur > ans {
					ans = cur
				}
			}
		}
	}
	return ans
}

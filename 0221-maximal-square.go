package leetCode

/*
https://leetcode.com/problems/maximal-square/
Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

Example:

Input:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

Output: 4
*/

// Approach 1: Brute Force
/*
The simplest approach consists of trying to find out every possible square of 1’s that can be formed from within
the matrix. The question now is – how to go for it?

We use a variable to contain the size of the largest square found so far and another variable to store the size of
the current, both initialized to 0. Starting from the left uppermost point in the matrix, we search for a 1.
No operation needs to be done for a 0. Whenever a 1 is found, we try to find out the largest square that can be
formed including that 1. For this, we move diagonally (right and downwards), i.e. we increment the row index and
column index temporarily and then check whether all the elements of that row and column are 1 or not. If all the
elements happen to be 1, we move diagonally further as previously. If even one element turns out to be 0, we stop
this diagonal movement and update the size of the largest square. Now we, continue the traversal of the matrix
from the element next to the initial 1 found, till all the elements of the matrix have been traversed.
*/

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	resp := 0
	rows, cols := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				sqlen := 1
				flag := true
				for sqlen+i < rows && sqlen+j < cols && flag {
					for k := j; k <= sqlen+j; k++ {
						if matrix[i+sqlen][k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						for k := i; k <= sqlen+i; k++ {
							if matrix[k][j+sqlen] == '0' {
								flag = false
								break
							}
						}
					}

					if flag {
						sqlen++
					}
				}
				if sqlen > resp {
					resp = sqlen
				}
			}
		}
	}

	return resp * resp
}

/* Approach 2: Dynamic Programming
We will explain this approach with the help of an example.

0 1 1 1 0
1 1 1 1 1
0 1 1 1 1
0 1 1 1 1
0 0 1 1 1
We initialize another matrix (dp) with the same dimensions as the original one initialized with all 0’s.

dp(i,j) represents the side length of the maximum square whose bottom right corner is the cell with index (i,j) in the original matrix.

Starting from index (0,0), for every 1 found in the original matrix, we update the value of the current element as

dp(i,j)=min(dp(i−1,j−1),dp(i−1,j),dp(i,j−1))+1.
*/
/*
[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
*/
func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}

	return n2
}

func maximalSquare2(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		dp[i] = make([]int, cols+1)
	}

	resp := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				if dp[i][j] > resp {
					resp = dp[i][j]
				}
			}
		}
	}

	return resp * resp
}

/* Approach 3: Better Dynamic Programming
In the previous approach for calculating dp of i^{th} row we are using only the previous element and the
(i-1)^{th}row. Therefore, we don't need 2D dp matrix as 1D dp array will be sufficient for this.

Initially the dp array contains all 0's. As we scan the elements of the original matrix across a row,
we keep on updating the dp array as per the equation dp[j]=min(dp[j−1],dp[j],prev), where prev refers to the old
dp[j−1]. For every row, we repeat the same process and update in the same dp array.
*/

func maximalSquare3(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	dp := make([]int, cols+1)
	resp, prev := 0, 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			tmp := dp[j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = min(prev, min(dp[j-1], dp[j])) + 1
				if dp[j] > resp {
					resp = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prev = tmp
		}
	}

	return resp * resp
}

package dp

/**
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] += grid[i-1][j]
		}
	}

	return grid[m-1][n-1]
}

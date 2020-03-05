package dp

/**
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?



An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:

Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid[0]), len(obstacleGrid)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			temp[i] = 1
		} else {
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				if j != 0 {
					temp[j] += temp[j-1]
				}
			} else {
				temp[j] = 0
			}
		}
	}
	return temp[n-1]
}

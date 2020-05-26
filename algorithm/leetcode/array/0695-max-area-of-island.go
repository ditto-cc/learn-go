package array

/**
 * Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
 *
 * Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)
 *
 * Example 1:
 *
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 *  [0,0,0,0,0,0,0,1,1,1,0,0,0],
 *  [0,1,1,0,1,0,0,0,0,0,0,0,0],
 *  [0,1,0,0,1,1,0,0,1,0,1,0,0],
 *  [0,1,0,0,1,1,0,0,1,1,1,0,0],
 *  [0,0,0,0,0,0,0,0,0,0,1,0,0],
 *  [0,0,0,0,0,0,0,1,1,1,0,0,0],
 *  [0,0,0,0,0,0,0,1,1,0,0,0,0]]
 * Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
 * Example 2:
 *
 * [[0,0,0,0,0,0,0,0]]
 * Given the above grid, return 0.
 * Note: The length of each dimension in the given grid does not exceed 50.
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/max-area-of-island
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func maxAreaOfIslandSolution(grid [][]int, i, j, m, n int, area *int) {
	if grid[i][j] == 0 {
		return
	}
	(*area)++
	grid[i][j] = 0
	if i > 0 {
		maxAreaOfIslandSolution(grid, i-1, j, m, n, area)
	}
	if j > 0 {
		maxAreaOfIslandSolution(grid, i, j-1, m, n, area)
	}
	if i < m {
		maxAreaOfIslandSolution(grid, i+1, j, m, n, area)
	}
	if j < n {
		maxAreaOfIslandSolution(grid, i, j+1, m, n, area)
	}
}

func maxAreaOfIsland(grid [][]int) int {

	m, n := len(grid)-1, len(grid[0])-1
	res, cur := 0, 0

	for i, row := range grid {
		for j := range row {

			cur = 0
			maxAreaOfIslandSolution(grid, i, j, m, n, &cur)
			if cur > res {
				res = cur
			}
		}
	}

	return res
}

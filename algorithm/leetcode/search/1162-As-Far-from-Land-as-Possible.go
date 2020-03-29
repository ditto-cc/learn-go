package search

/**
 * Given an N x N grid containing only values 0 and 1, where 0 represents water and 1 represents land, find a water cell such that its distance to the nearest land cell is maximized and return the distance.
 *
 * The distance used in this problem is the Manhattan distance: the distance between two cells (x0, y0) and (x1, y1) is |x0 - x1| + |y0 - y1|.
 *
 * If no land or water exists in the grid, return -1.
 *
 *
 *
 * Example 1:
 *
 *
 *
 * Input: [[1,0,1],[0,0,0],[1,0,1]]
 * Output: 2
 * Explanation:
 * The cell (1, 1) is as far as possible from all the land with distance 2.
 * Example 2:
 *
 *
 *
 * Input: [[1,0,0],[0,0,0],[0,0,0]]
 * Output: 4
 * Explanation:
 * The cell (2, 2) is as far as possible from all the land with distance 4.
 *
 *
 * Note:
 *
 * 1 <= grid.length == grid[0].length <= 100
 * grid[i][j] is 0 or 1
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/as-far-from-land-as-possible
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func maxDistance(grid [][]int) int {
	n := len(grid)
	q := make([]int, 0)
	for i, row := range grid {
		for j, e := range row {
			if e == 1 {
				q = append(q, i*n+j)
			}
		}
	}

	levelNum := len(q)
	if levelNum == n*n || levelNum == 0 {
		return -1
	}

	res := -1
	for ; levelNum > 0; levelNum = len(q) {
		for i := 0; i < levelNum; i++ {
			x, y := q[i]/n, q[i]%n
			if x-1 >= 0 && grid[x-1][y] == 0 {
				grid[x-1][y] = 1
				q = append(q, q[i]-n)
			}

			if x+1 < n && grid[x+1][y] == 0 {
				grid[x+1][y] = 1
				q = append(q, q[i]+n)
			}

			if y-1 >= 0 && grid[x][y-1] == 0 {
				grid[x][y-1] = 1
				q = append(q, q[i]-1)
			}

			if y+1 < n && grid[x][y+1] == 0 {
				grid[x][y+1] = 1
				q = append(q, q[i]+1)
			}
		}
		q = q[levelNum:]
		res++
	}
	return res
}

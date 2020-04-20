package search

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

示例 1:

输入:
11110
11010
11000
00000
输出: 1
示例 2:

输入:
11000
11000
00100
00011
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	var res = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				floodFill(grid, i, j, m, n)
			}
		}
	}
	return res
}

func floodFill(grid [][]byte, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	floodFill(grid, i-1, j, m, n)
	floodFill(grid, i+1, j, m, n)
	floodFill(grid, i, j-1, m, n)
	floodFill(grid, i, j+1, m, n)
}

package array

/**
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

Example 1:

Input:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
Output:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
Example 2:

Input:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
Output:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
Follow up:

A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/set-matrix-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func setZeroes(matrix [][]int) {
	rowSet, colSet := map[int]int{}, map[int]int{}
	for i, row := range matrix {
		for j, cell := range row {
			if cell != 0 {
				continue
			}
			rowSet[i] = 0
			colSet[j] = 0
		}
	}

	for _, row := range rowSet {
		for i := range matrix[row] {
			matrix[row][i] = 0
		}
	}

	m := len(matrix)
	for _, col := range colSet {
		for i := 0; i < m; i++ {
			matrix[i][col] = 0
		}
	}
}

//
//func setZeroes(matrix [][]int) {
//	m, n := len(matrix), len(matrix[0])
//	firstRowZero := false
//
//	for _, e := range matrix[0] {
//		if e == 0 {
//			firstRowZero = true
//			break
//		}
//	}
//
//	for i := 1; i < m; i++ {
//		for j, cell := range matrix[i] {
//			if cell == 0 {
//				matrix[0][j] = 0
//				matrix[i][0] = 0
//			}
//		}
//	}
//
//	for i := 1; i < m; i++ {
//		for j := n - 1; j >= 0; j-- {
//			if matrix[0][j] == 0 || matrix[i][0] == 0 {
//				matrix[i][j] = 0
//			}
//		}
//	}
//
//	if !firstRowZero {
//		return
//	}
//
//	for i, _ := range matrix[0] {
//		matrix[0][i] = 0
//	}
//}

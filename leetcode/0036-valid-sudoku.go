package leetcode

/**
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

A partially filled sudoku which is valid.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Example 1:

Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
Example 2:

Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9 and the character '.'.
The given board size is always 9x9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numToIndex(e byte) int {
	switch e {
	case '1':
		return 0
	case '2':
		return 1
	case '3':
		return 2
	case '4':
		return 3
	case '5':
		return 4
	case '6':
		return 5
	case '7':
		return 6
	case '8':
		return 7
	case '9':
		return 8
	default:
		return -1
	}
}

func isValidSudoku(board [][]byte) bool {
	rowMemo := make([][]bool, 9, 9)
	colMemo := make([][]bool, 9, 9)
	for i := 0; i < 9; i++ {
		rowMemo[i], colMemo[i] = make([]bool, 9, 9), make([]bool, 9, 9)
	}

	startPos := [][]int{{0, 0}, {0, 3}, {0, 6}, {3, 0}, {3, 3}, {3, 6}, {6, 0}, {6, 3}, {6, 6}}
	boxMemo := make([]bool, 9, 9)

	for _, pos := range startPos {
		for i := range boxMemo {
			boxMemo[i] = false
		}
		for i, limitI := pos[0], pos[0]+3; i < limitI; i++ {
			for j, limitJ := pos[1], pos[1]+3; j < limitJ; j++ {
				if index := numToIndex(board[i][j]); index == -1 {
					continue
				} else {
					if boxMemo[index] || rowMemo[i][index] || colMemo[j][index] {
						return false
					}
					boxMemo[index], rowMemo[i][index], colMemo[j][index] = true, true, true
				}
			}
		}
	}

	return true
}

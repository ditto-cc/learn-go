package tree

/**
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1
Return:

[
   [5,4,11,2],
   [5,8,4,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	pathSumSolution(root, &res, &[]int{}, sum)
	return res
}

func pathSumSolution(root *TreeNode, res *[][]int, path *[]int, sum int) {
	if root == nil {
		return
	}
	remain := sum - root.Val
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil && remain == 0 {
		temp := make([]int, len(*path), len(*path))
		copy(temp, *path)
		*res = append(*res, temp)
	}
	pathSumSolution(root.Left, res, path, remain)
	pathSumSolution(root.Right, res, path, remain)

	*path = (*path)[:len(*path)-1]
}

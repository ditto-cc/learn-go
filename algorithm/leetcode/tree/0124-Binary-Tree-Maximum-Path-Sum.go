package tree

import "math"

/**
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6
Example 2:

Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
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

func maxPath(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	lp, rp := maxInt(maxPath(root.Left, max), 0), maxInt(maxPath(root.Right, max), 0)
	if root.Val+lp+rp > *max {
		*max = root.Val + lp + rp
	}

	return maxInt(lp, rp) + root.Val
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	maxPath(root, &res)
	return res
}

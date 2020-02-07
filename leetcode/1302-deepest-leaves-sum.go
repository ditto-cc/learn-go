package leetcode

import "math"

/**
Given a binary tree, return the sum of values of its deepest leaves.


Example 1:



Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
Output: 15


Constraints:

The number of nodes in the tree is between 1 and 10^4.
The value of nodes is between 1 and 100.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/deepest-leaves-sum
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

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(height(node.Left)), float64(height(node.Right))))
}

func countDeepestLeaves(node *TreeNode, height int) int {
	if node == nil {
		return 0
	}
	res := 0
	if 1 == height {
		res += node.Val
	}
	return res + countDeepestLeaves(node.Left, height-1) + countDeepestLeaves(node.Right, height-1)
}

func deepestLeavesSum(root *TreeNode) int {
	h := height(root)
	return countDeepestLeaves(root, h)
}

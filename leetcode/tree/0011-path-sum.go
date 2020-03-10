package tree

/**
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	remain := sum - root.Val
	switch {
	case root.Left == nil && root.Right == nil:
		return remain == 0
	case root.Left == nil:
		return hasPathSum(root.Right, remain)
	case root.Right == nil:
		return hasPathSum(root.Left, remain)
	default:
		return hasPathSum(root.Left, remain) || hasPathSum(root.Right, remain)
	}
}

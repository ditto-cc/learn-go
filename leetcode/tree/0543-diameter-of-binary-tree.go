package tree

/**
Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
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

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func diameterOfBinTreeSolution(node *TreeNode, dia *int) int {
	if node == nil {
		return 0
	}
	leftH := diameterOfBinTreeSolution(node.Left, dia)
	rightH := diameterOfBinTreeSolution(node.Right, dia)
	*dia = max(leftH+rightH, *dia)
	return 1 + max(leftH, rightH)
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	diameterOfBinTreeSolution(root, &res)
	return res
}

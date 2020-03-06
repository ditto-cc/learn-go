package tree

import "learn-go/leetcode"

/**
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
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
func inorderTraversal(root *leetcode.TreeNode) []int {
	s := []*leetcode.TreeNode{}
	res := []int{}
	p := root
	var top *leetcode.TreeNode
	for len(s) > 0 || p != nil {
		if p != nil {
			s = append(s, p)
			p = p.Left
		} else {
			topIndex := len(s) - 1
			top = s[topIndex]
			res = append(res, top.Val)
			s = s[:topIndex]
			p = top.Right
		}
	}
	return res
}

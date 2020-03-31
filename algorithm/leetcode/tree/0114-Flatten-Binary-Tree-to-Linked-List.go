package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var pre *TreeNode = nil
	s := []*TreeNode{root}
	for len(s) > 0 {
		top := s[0]
		s = s[:len(s)-1]
		if pre != nil {
			pre.Left = nil
			pre.Right = top
		}

		if top.Right != nil {
			s = append(s, top.Right)
		}

		if top.Left != nil {
			s = append(s, top.Left)
		}
		pre = top
	}
}

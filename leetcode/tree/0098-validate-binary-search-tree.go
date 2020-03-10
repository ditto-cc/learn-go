package tree

import "math"

/**
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:

    2
   / \
  1   3

Input: [2,1,3]
Output: true
Example 2:

    5
   / \
  1   4
     / \
    3   6

Input: [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
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

func isValidBST(root *TreeNode) bool {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	s := []*TreeNode{}
	p := root
	pre := math.Inf(-1)
	for p != nil || len(s) > 0 {
		if p != nil {
			s = append(s, p)
			p = p.Left
		} else {
			top := len(s) - 1
			cur := s[top].Val
			if float64(cur) <= pre {
				return false
			}
			pre = float64(cur)
			p = s[top].Right
			s = s[:top]
		}
	}

	return true
}

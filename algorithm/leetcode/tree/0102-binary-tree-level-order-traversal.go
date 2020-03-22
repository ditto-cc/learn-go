package tree

/**
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
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
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{}
	q = append(q, root)
	var frontNode *TreeNode
	for len(q) > 0 {
		levelNodeNum := len(q)
		levelArr := []int{}
		for ; levelNodeNum > 0; levelNodeNum-- {
			frontNode = q[0]
			q = q[1:]
			levelArr = append(levelArr, frontNode.Val)
			if frontNode.Left != nil {
				q = append(q, frontNode.Left)
			}
			if frontNode.Right != nil {
				q = append(q, frontNode.Right)
			}
		}
		res = append(res, levelArr)
	}

	return res
}

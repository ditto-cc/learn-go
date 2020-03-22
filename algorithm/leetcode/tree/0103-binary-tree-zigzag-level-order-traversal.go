package tree

/**
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	nodeNum, nextNum := 1, 0
	reverse := false
	for nodeNum > 0 {
		level := make([]int, nodeNum, nodeNum)
		for i := 0; i < nodeNum; i++ {
			node := q[i]
			if reverse {
				level[nodeNum-i-1] = node.Val
			} else {
				level[i] = node.Val
			}
			l, r := node.Left, node.Right
			if l != nil {
				q = append(q, l)
				nextNum++
			}
			if r != nil {
				q = append(q, r)
				nextNum++
			}
		}
		reverse = !reverse
		q = q[nodeNum:]
		nodeNum = nextNum
		nextNum = 0
		res = append(res, level)
	}

	return res
}

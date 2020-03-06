package tree

import "learn-go/leetcode"

/**
Given a binary tree, return the sum of values of nodes with even-valued grandparent.  (A grandparent of a node is the parent of its parent, if it exists.)

If there are no nodes with an even-valued grandparent, return 0.



Example 1:



Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 18
Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.


Constraints:

The number of nodes in the tree is between 1 and 10^4.
The value of nodes is between 1 and 100.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent
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
func sumOfNode(node *leetcode.TreeNode) int {
	if node == nil || node.Val%2 == 1 {
		return 0
	}
	res := 0
	left, right := node.Left, node.Right
	if left != nil {
		ll, lr := left.Left, left.Right
		if ll != nil {
			res += ll.Val
		}
		if lr != nil {
			res += lr.Val
		}
	}
	if right != nil {
		rl, rr := right.Left, right.Right
		if rl != nil {
			res += rl.Val
		}
		if rr != nil {
			res += rr.Val
		}
	}
	return res
}

func sumEvenGrandparent(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}
	res := sumOfNode(root)
	res += sumEvenGrandparent(root.Left) + sumEvenGrandparent(root.Right)
	return res
}

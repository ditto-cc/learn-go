package tree

import "learn-go/leetcode"

/**
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
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

func findElementIndex(arr []int, l, r, val int) int {
	index := -1
	for i := l; i <= r; i++ {
		if arr[i] == val {
			index = i
			break
		}
	}
	return index
}

func buildTreeNode(pre, in []int, pl, pr, il, ir int) *leetcode.TreeNode {
	if pl > pr {
		return nil
	}
	node := &leetcode.TreeNode{Val: pre[pl]}
	imid := findElementIndex(in, il, ir, pre[pl])
	node.Left = buildTreeNode(pre, in, pl+1, pl+imid-il, il, imid-1)
	node.Right = buildTreeNode(pre, in, pl+imid-il+1, pr, imid+1, ir)
	return node
}

func buildTree(preorder []int, inorder []int) *leetcode.TreeNode {
	return buildTreeNode(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

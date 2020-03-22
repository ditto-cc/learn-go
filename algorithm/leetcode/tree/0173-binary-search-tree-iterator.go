package tree

/**
Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.



Example:



BSTIterator iterator = new BSTIterator(root);
iterator.next();    // return 3
iterator.next();    // return 7
iterator.hasNext(); // return true
iterator.next();    // return 9
iterator.hasNext(); // return true
iterator.next();    // return 15
iterator.hasNext(); // return true
iterator.next();    // return 20
iterator.hasNext(); // return false


Note:

next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-search-tree-iterator
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
type BSTIterator struct {
	s []*TreeNode
	p *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{s: []*TreeNode{}, p: root}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	var ret int
	for len(this.s) > 0 || this.p != nil {
		if this.p != nil {
			this.s = append(this.s, this.p)
			this.p = this.p.Left
		} else {
			top := this.s[len(this.s)-1]
			ret = top.Val
			this.s = this.s[:len(this.s)-1]
			this.p = top.Right
			break
		}
	}
	return ret
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.s) > 0 || this.p != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

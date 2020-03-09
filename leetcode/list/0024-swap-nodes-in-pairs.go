package list

/**
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.



Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p, c, n *ListNode
	p, c = head, head.Next
	p.Next = c.Next
	c.Next = p
	head = c

	for p.Next != nil && p.Next.Next != nil {
		c, n = p.Next, p.Next.Next

		c.Next = n.Next
		n.Next = c
		p.Next = n

		p = c
	}

	return head
}

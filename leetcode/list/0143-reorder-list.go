package list

import "learn-go/leetcode"

/**
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:

Given 1->2->3->4, reorder it to 1->4->2->3.
Example 2:

Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middle(head, tail *leetcode.ListNode) *leetcode.ListNode {
	p, q := head, head
	for q.Next != tail && q.Next.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	return p
}

func reorderList(head *leetcode.ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := middle(head, nil)
	var pre, cur, next *leetcode.ListNode
	pre, cur = nil, mid.Next
	mid.Next = nil
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	p, q := head, pre
	for q != nil {
		pn, qn := p.Next, q.Next
		q.Next, p.Next = pn, q
		p, q = pn, qn
	}
}

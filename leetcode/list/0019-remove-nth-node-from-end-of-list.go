package list

import "learn-go/leetcode"

/**
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *leetcode.ListNode, n int) *leetcode.ListNode {
	if head == nil {
		return head
	}
	dummy := &leetcode.ListNode{Next: head}
	p, q := dummy, dummy
	for ; n > 0; n-- {
		q = q.Next
	}
	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}

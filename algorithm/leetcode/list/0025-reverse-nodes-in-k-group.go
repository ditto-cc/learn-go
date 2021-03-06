package list

/**
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {

	dummy := &ListNode{Next: head}
	i := 0
	for p := dummy; p.Next != nil; {
		q := p.Next
		for i = 0; i < k && q != nil; i, q = i+1, q.Next {

		}

		if i < k {
			break
		}

		pre, cur := p, p.Next
		for i = 0; i < k && cur != q; i++ {
			next := cur.Next
			cur.Next = pre
			pre, cur = cur, next
		}

		tail := p.Next
		tail.Next = cur
		p.Next = pre
		p = tail
	}

	return dummy.Next
}

package list

/**
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head = &ListNode{Next: head}
	pre, cur, next := head, head.Next, head.Next.Next
	for cur != nil && next != nil {
		if cur.Val != next.Val {
			if pre.Next != cur {
				pre.Next = next
			} else {
				pre = cur
			}
		}
		cur = next
		next = next.Next
	}
	if pre.Next != cur {
		pre.Next = next
	}

	return head.Next
}

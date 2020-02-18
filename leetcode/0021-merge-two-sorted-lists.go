package leetcode

/**
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{Next:l1}
	l1 = head
	for ; l1.Next != nil && l2 != nil; l1 = l1.Next {
		if l1.Next.Val > l2.Val {
			n2 := l2.Next
			l2.Next = l1.Next
			l1.Next = l2
			l2 = n2
		}
	}

	if l1.Next == nil {
		l1.Next = l2
	}

	return head.Next
}

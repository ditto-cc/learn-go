package leetcode

/**
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	head := l1
	l1, l2 = &ListNode{Next:l1}, &ListNode{Next:l2}

	for ; l1.Next != nil && l2.Next != nil; l1, l2 = l1.Next, l2.Next {
		l1.Next.Val += l2.Next.Val + carry
		carry = l1.Next.Val / 10
		l1.Next.Val %= 10
	}

	if l1.Next == nil {
		l1.Next = l2.Next
	}

	for ; carry != 0; l1 = l1.Next {
		if l1.Next == nil {
			l1.Next = &ListNode{Val: carry}
			break
		}
		l1.Next.Val += carry
		carry = l1.Next.Val / 10
		l1.Next.Val %= 10
	}

	return head
}
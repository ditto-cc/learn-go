package list

/**
You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

Example:

Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
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
	p, q := l1, l2
	var s1, s2, s3 []*ListNode
	for i := 0; i < 2; {
		p = p.Next
		q = q.Next
		if p == nil {
			p = l2
			i++
		}
		if q == nil {
			q = l1
			i++
		}
	}

	if q != l1 {
		l1, l2 = l2, l1
		p, q = q, p
	}

	for i := p; i != nil; {
		s2 = append(s2, i)
		s3 = append(s3, q)
		i = i.Next
		q = q.Next
	}

	carry := 0
	for i := len(s2) - 1; i >= 0; i-- {
		s2[i].Val += s3[i].Val + carry
		carry = s2[i].Val / 10
		s2[i].Val %= 10
	}

	if carry == 0 {
		return l2
	}

	for i := l2; i != p; i = i.Next {
		s1 = append(s1, i)
	}

	for i := len(s1) - 1; i >= 0 && carry > 0; i-- {
		s1[i].Val += carry
		carry = s1[i].Val / 10
		s1[i].Val %= 10
	}

	if carry == 0 {
		return l2
	}
	return &ListNode{Val: carry, Next: l2}
}

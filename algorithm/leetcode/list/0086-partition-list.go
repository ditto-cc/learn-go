package list

/**
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:

Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head}

	q := dummy
	for ; q.Next != nil && q.Next.Val < x; q = q.Next {

	}

	if q.Next == nil {
		return head
	}

	for p := q.Next; p.Next != nil; {
		if p.Next.Val < x {
			node := p.Next
			p.Next = node.Next

			node.Next = q.Next
			q.Next = node
			q = node
		} else {
			p = p.Next
		}
	}

	return dummy.Next
}

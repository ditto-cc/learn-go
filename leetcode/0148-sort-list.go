package leetcode

/**
Sort a linked list in O(n log n) time using constant space complexity.

Example 1:

Input: 4->2->1->3
Output: 1->2->3->4
Example 2:

Input: -1->5->3->4->0
Output: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middle(head, tail *ListNode) *ListNode {
	p, q := head, head
	for q.Next != tail && q.Next.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	return p
}

func merge(head, mid, tail *ListNode) {

	p, q := head, mid
	for p.Next != mid.Next || q.Next != tail {
		if p.Next == mid.Next || q.Next == tail {
			break
		} else if p.Next.Val < q.Next.Val {
			p = p.Next
		} else {
			del := q.Next
			q.Next = del.Next
			del.Next = p.Next
			p.Next = del
			p = del
		}
	}
}

func mergeSort(head, tail *ListNode)  {
	if head == tail || head.Next == tail || head.Next.Next == tail {
		return
	}
	mid := middle(head, tail)
	mergeSort(head, mid.Next)
	mergeSort(mid, tail)
	merge(head, mid, tail)
}

func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{Next:head}
	mergeSort(dummy, nil)
	return dummy.Next
}
package list

import "learn-go/leetcode"

/**
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func down(h []*leetcode.ListNode, i, end int) {
	for i*2+1 < end {
		lc := i*2 + 1
		if lc+1 < end && h[lc+1].Val < h[lc].Val {
			lc++
		}
		if h[i].Val < h[lc].Val {
			break
		}
		h[i], h[lc] = h[lc], h[i]
		i = lc
	}
}

func mergeKLists(lists []*leetcode.ListNode) *leetcode.ListNode {
	for i := 0; i < len(lists); {
		if lists[i] == nil {
			lists = append(lists[:i], lists[i+1:]...)
		} else {
			i++
		}
	}
	n := len(lists)
	if n == 0 {
		return nil
	} else if n < 2 {
		return lists[0]
	}

	dummy := &leetcode.ListNode{}
	p := dummy
	for i := (n - 1) / 2; i >= 0; i-- {
		down(lists, i, n)
	}

	for n > 0 {
		p.Next = lists[0]
		p = p.Next
		if lists[0].Next != nil {
			lists[0] = lists[0].Next
		} else {
			n--
			lists[0] = lists[n]
			lists = lists[:n]
		}
		down(lists, 0, n)
	}

	return dummy.Next
}

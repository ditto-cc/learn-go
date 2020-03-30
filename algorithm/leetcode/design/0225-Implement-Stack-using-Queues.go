package design

/**
 * Implement the following operations of a stack using queues.
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * empty() -- Return whether the stack is empty.
 * Example:
 *
 * MyStack stack = new MyStack();
 *
 * stack.push(1);
 * stack.push(2);
 * stack.top();   // returns 2
 * stack.pop();   // returns 2
 * stack.empty(); // returns false
 * Notes:
 *
 * You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.
 * Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
 * You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/implement-stack-using-queues
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type MyStack struct {
	q1, q2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q1: make([]int, 0),
		q2: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if len(this.q1) > 0 {
		this.q1 = append(this.q1, x)
	} else {
		this.q2 = append(this.q2, x)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.q1) > 0 {
		ret := this.q1[len(this.q1)-1]
		this.q1 = this.q1[:len(this.q1)-1]
		this.q2 = this.q1
		this.q1 = make([]int, 0)
		return ret
	} else if len(this.q2) > 0 {
		ret := this.q2[len(this.q2)-1]
		this.q2 = this.q2[:len(this.q2)-1]
		this.q1 = this.q2
		this.q2 = make([]int, 0)
		return ret
	}
	return -1
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.q1) > 0 {
		return this.q1[len(this.q1)-1]
	} else if len(this.q2) > 0 {
		return this.q2[len(this.q2)-1]
	}
	return -1
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q1) == 0 && len(this.q2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

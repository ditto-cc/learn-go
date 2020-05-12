package stack

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。


示例:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.


提示：

pop、top 和 getMin 操作总是在 非空栈 上调用。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MinStack struct {
	data, s      []int
	dSize, sSize int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{dSize: -1, sSize: -1}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	this.dSize++
	if this.sSize == -1 || this.s[this.sSize] >= x {
		this.s = append(this.s, x)
		this.sSize++
	}
}

func (this *MinStack) Pop() {
	top := this.Top()
	this.data = this.data[:this.dSize]
	this.dSize--
	if top == this.s[this.sSize] {
		this.s = this.s[:this.sSize]
		this.sSize--
	}
}

func (this *MinStack) Top() int {
	if this.dSize == -1 {
		panic("Empty Stack")
	}
	return this.data[this.dSize]
}

func (this *MinStack) GetMin() int {
	if this.dSize == -1 {
		panic("Empty Stack")
	}
	return this.s[this.sSize]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

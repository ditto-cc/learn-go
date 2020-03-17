package design

/**
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2 /capacity / );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type myListNode struct {
	key, val    int
	prior, next *myListNode
}

func (this *LRUCache) moveForward(node *myListNode) {
	node.prior.next = node.next
	node.next.prior = node.prior

	node.next = this.head.next
	node.prior = this.head

	this.head.next.prior = node
	this.head.next = node
}

type LRUCache struct {
	capacity int
	table    map[int]*myListNode
	head     *myListNode
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{
		capacity: capacity,
		table:    map[int]*myListNode{},
		head:     &myListNode{},
	}
	ret.head.next, ret.head.prior = ret.head, ret.head
	return ret
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.table[key]; ok {
		this.moveForward(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.table[key]; ok {
		node.val = value
		this.moveForward(node)
	} else {
		var newNode *myListNode
		if len(this.table) == this.capacity {
			newNode = this.head.prior
			delete(this.table, newNode.key)
			newNode.key, newNode.val = key, value
			this.moveForward(newNode)
		} else {
			newNode = &myListNode{
				prior: this.head,
				next:  this.head.next,
				key:   key,
				val:   value,
			}
			this.head.next.prior = newNode
			this.head.next = newNode
		}
		this.table[key] = newNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

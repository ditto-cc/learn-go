package design

/**
 * Design and implement a data structure for Least Frequently Used (LFU) cache. It should support the following operations: get and put.
 *
 * get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least frequently used item before inserting a new item. For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency), the least recently used key would be evicted.
 *
 * Note that the number of times an item is used is the number of calls to the get and put functions for that item since it was inserted. This number is set to zero when the item is removed.
 *
 *
 *
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 *
 *
 *
 * Example:
 *
 * LFUCache cache = new LFUCache( 2 );
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // returns 1
 * cache.put(3, 3);    // evicts key 2
 * cache.get(2);       // returns -1 (not found)
 * cache.get(3);       // returns 3.
 * cache.put(4, 4);    // evicts key 1.
 * cache.get(1);       // returns -1 (not found)
 * cache.get(3);       // returns 3
 * cache.get(4);       // returns 4
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/lfu-cache
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type freqNode struct {
	key, value, fre int
	prior, next     *freqNode
}

type LFUCache struct {
	minFreq, capacity int
	freqMap, dataMap  map[int]*freqNode
}

func newFreqNode(key, val int) *freqNode {
	return &freqNode{key: key, value: val, fre: 1}
}

func (this *LFUCache) update(node *freqNode) {
	if node.fre == this.minFreq && node.next.next == node {
		this.minFreq++
	}
	node.fre++
	node.prior.next, node.next.prior = node.next, node.prior
	this.addToFreqMap(node)
}

func (this *LFUCache) addToFreqMap(node *freqNode) {
	if head, ok := this.freqMap[node.fre]; ok {
		node.prior, node.next = head, head.next
		head.next.prior, head.next = node, node
	} else {
		head = newFreqNode(-1, -1)
		node.prior, node.next = head, head
		head.next, head.prior = node, node
		this.freqMap[node.fre] = head
	}
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		minFreq:  0,
		freqMap:  map[int]*freqNode{},
		dataMap:  map[int]*freqNode{},
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.dataMap[key]; ok {
		this.update(node)
		return node.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}

	if node, ok := this.dataMap[key]; ok {
		node.value = value
		this.update(node)
	} else {
		node = newFreqNode(key, value)
		if len(this.dataMap) == this.capacity {
			node := this.freqMap[this.minFreq].prior
			node.prior.next, node.next.prior = node.next, node.prior
			delete(this.dataMap, node.key)
		}
		this.minFreq = 1
		this.addToFreqMap(node)
		this.dataMap[key] = node
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

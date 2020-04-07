package main

import "fmt"

func main() {
	cache := Constructor(0)
	cache.Put(1, 1)
	fmt.Println(cache.Get(1))
	cache.Put(2, 2)
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
}

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

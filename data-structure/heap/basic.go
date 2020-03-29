package heap

import (
	"learn-go/data-structure/utils"
)

type Heap struct {
	data []utils.Comparable
	size int
}

func CreateHeap() *Heap {
	return &Heap{data: []utils.Comparable{}}
}

func (h *Heap) Push(e utils.Comparable) {
	h.data = append(h.data, e)
	h.size++
	h.up(h.Size() - 1)
}

func (h *Heap) Pop() utils.Comparable {
	if h.Size() == 0 {
		return nil
	}
	n := h.Size() - 1
	ret := h.First()
	h.data[0] = h.data[n]
	h.down(0, n)
	h.data = h.data[:n]
	h.size--
	return ret
}

func (h *Heap) First() utils.Comparable {
	if h.Size() == 0 {
		return nil
	}
	return h.data[0]
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Heapify() {
	for i, n := (h.Size()-1)/2, h.Size(); i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *Heap) down(i, end int) {
	var c int
	for 2*i+1 < end {
		c = 2*i + 1
		if c+1 < end && h.data[c+1].Compare(h.data[c]) < 0 {
			c++
		}

		if h.data[i].Compare(h.data[c]) <= 0 {
			break
		}
		h.swap(i, c)
		i = c
	}
}

func (h *Heap) up(i int) {
	var p int
	for i > 0 {
		p = (i - 1) / 2
		if h.data[i].Compare(h.data[p]) >= 0 {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

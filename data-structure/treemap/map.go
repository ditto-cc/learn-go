package treemap

import (
	"learn-go/data-structure/compare"
	"learn-go/data-structure/tree/avl"
)

type TreeMap struct {
	tree *avl.AVLTree
}

func CreateTreeMap() *TreeMap {
	return &TreeMap{tree: avl.CreateAVL()}
}

func (m *TreeMap) Get(key compare.Comparable) interface{} {
	return m.tree.Get(key)
}

func (m *TreeMap) Set(key compare.Comparable, val interface{}) bool {
	return m.tree.Set(key, val)
}

func (m *TreeMap) Add(key compare.Comparable, val interface{}) {
	m.tree.Add(key, val)
}

func (m *TreeMap) Contains(key compare.Comparable) bool {
	return m.tree.Contains(key)
}

func (m *TreeMap) Remove(key compare.Comparable) interface{} {
	return m.tree.Remove(key)
}

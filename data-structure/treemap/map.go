package treemap

import (
	"learn-go/data-structure/tree/avl"
	"learn-go/data-structure/utils"
)

type TreeMap struct {
	tree *avl.AVLTree
}

func CreateTreeMap() *TreeMap {
	return &TreeMap{tree: avl.CreateAVL()}
}

func (m *TreeMap) Get(key utils.Comparable) interface{} {
	return m.tree.Get(key)
}

func (m *TreeMap) Set(key utils.Comparable, val interface{}) bool {
	return m.tree.Set(key, val)
}

func (m *TreeMap) Add(key utils.Comparable, val interface{}) {
	m.tree.Add(key, val)
}

func (m *TreeMap) Contains(key utils.Comparable) bool {
	return m.tree.Contains(key)
}

func (m *TreeMap) Remove(key utils.Comparable) interface{} {
	return m.tree.Remove(key)
}

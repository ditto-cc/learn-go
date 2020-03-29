package entity

import (
	"fmt"
	"learn-go/data-structure/tree/avl"
	"learn-go/data-structure/tree/bst"
)

type Student struct {
	Name  string
	Score float64
}

func (s *Student) String() string {
	return fmt.Sprintf("(%s, %v)", s.Name, s.Score)
}

func compare(a, b *Student) int {
	res := int(b.Score - a.Score)
	if res == 0 {
		if a.Name > b.Name {
			res = 1
		} else if a.Name < b.Name {
			res = -1
		}
	}
	return res
}

func (a *Student) CompareAVL(b avl.Comparable) int {
	if c, ok := b.(*Student); ok {
		return compare(a, c)
	} else {
		panic("Student Type Assertion Failed.")
	}
}

func (a *Student) CompareBST(b bst.Comparable) int {
	if c, ok := b.(*Student); ok {
		return compare(a, c)
	} else {
		panic("Student Type Assertion Failed.")
	}
}

package entity

import (
	"fmt"
	"learn-go/data-structure/bst"
)

type Student struct {
	Name  string
	Score float64
}

func (s *Student) String() string {
	return fmt.Sprintf("(%s, %v)", s.Name, s.Score)
}

func (a *Student) Compare(b bst.Comparable) int {
	c := b.(*Student)
	if a.Score < c.Score {
		return -1
	} else if a.Score > c.Score {
		return 1
	}
	return 0
}

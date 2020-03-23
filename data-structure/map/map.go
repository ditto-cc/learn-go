package _map

import (
	"learn-go/data-structure/compare"
)

type Map interface {
	Add(compare.Comparable, interface{})
	Remove(compare.Comparable) interface{}
	Contains(compare.Comparable) bool
	Set(compare.Comparable, interface{}) bool
}

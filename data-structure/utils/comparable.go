package utils

type Comparable interface {
	Compare(Comparable) int
}

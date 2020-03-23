package compare

type Comparable interface {
	Compare(Comparable) int
}

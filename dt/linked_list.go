package dt

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Float interface {
	~float64 | ~float32
}
type String interface {
	~string | ~byte
}

type Comparable interface {
	Integer | Float | String
}

type LinkedList[T Comparable] struct {
	val  T
	next *LinkedList[T]
}

func Traverse[T Comparable](head *LinkedList[T], k T) *LinkedList[T] {
	for node := head.next; node != nil; node = node.next {
		if node != nil && node.val == k {

		}
	}
}
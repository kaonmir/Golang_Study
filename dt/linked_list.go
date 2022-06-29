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

func MakeLinkedList[T Comparable](vals ...T) *LinkedList[T] {
	head := &LinkedList[T]{}
	node := head

	for _, val := range vals {
		node.next = &LinkedList[T]{val: val, next: nil}
		node = node.next
	}

	return head
}

/*
func Traverse[T Comparable](head *LinkedList[T], k T) *LinkedList[T] {
	for node := head.next; node != nil; node = node.next {
		if node != nil && node.val == k {
			return node
		}
	}

	return nil
}
*/

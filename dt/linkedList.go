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
	Val  T
	Next *LinkedList[T]
}

func MakeLinkedList[T Comparable](vals ...T) *LinkedList[T] {
	head := &LinkedList[T]{}
	node := head

	for _, val := range vals {
		node.Next = &LinkedList[T]{Val: val, Next: nil}
		node = node.Next
	}

	return head
}

func Traverse[T Comparable](head *LinkedList[T], k T) *LinkedList[T] {
	for node := head.Next; node != nil; node = node.Next {
		if node != nil && node.Val == k {
			return node
		}
	}
	return nil
}

func At[T Comparable](head *LinkedList[T], idx int) *LinkedList[T] {
	node := head
	for i := 0; i <= idx; i++ {
		node = node.Next
	}
	return node
}

func Insert[T Comparable](head *LinkedList[T], k T, idx int) {

}

package dt_test

import (
	"testing"

	"kaonmir.xyz/dt/dt"
)

func TestError(t *testing.T) {
	// t.Error("It is right to get an error")
}

func TestMakeLinkedList(t *testing.T) {
	head := dt.MakeLinkedList(1, 2)
	node := head.Next

	// Check value of the first node is 1
	if node == nil || node.Val != 1 {
		t.Error("The value of the first node is not 1")
	}

	node = node.Next
	if node == nil || node.Val != 2 {
		t.Error("The value of the second node is not 2")
	}
}

func TestTraverse(t *testing.T) {
	head := dt.MakeLinkedList(1, 2, 10, 11)

	// TestCase 1: traverse the first value 1
	if dt.Traverse(head, 1).Val != 1 {
		t.Error("The value of the first value is not 1")
	}

	// TestCase 2: traverse the thrid value 10
	if dt.Traverse(head, 10).Val != 10 {
		t.Error("The value of the thrid value is not 10")
	}

	// TestCase 3: traverse the worng value
	if dt.Traverse(head, 100) != nil {
		t.Error("The value of the wrong value is not nil")
	}
}

func TestAt(t *testing.T) {
	head := dt.MakeLinkedList(1, 2, 10, 11)

	// TestCase 1: The value of index 1 is 2
	if dt.At(head, 1).Val != 2 {
		t.Error("The value of index 1 is not 2")
	}

	// TestCase 2: The value of index 5 doesn't exist
	// TODO: I do not know how to implement this
}

package iterator

import (
	"testing"
)

func TestInOrderIterator(t *testing.T) {
	// Tree:    1
	//         / \
	//        2   3
	root := NewNode(1,
		NewTerminalNode(2),
		NewTerminalNode(3))
	tree := NewBinaryTree(root)
	iter := tree.InOrder()

	var got []int
	for iter.Next() {
		got = append(got, iter.Current.Value)
	}
	want := []int{2, 1, 3}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("in-order traversal mismatch at index %d: got %d, want %d", i, got[i], want[i])
		}
	}

	// Test Reset
	iter.Reset()
	got = got[:0]
	for iter.Next() {
		got = append(got, iter.Current.Value)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("[after reset] in-order traversal mismatch at index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}

func TestInOrderIterator_SingleNode(t *testing.T) {
	root := NewTerminalNode(42)
	tree := NewBinaryTree(root)
	iter := tree.InOrder()

	if !iter.Next() {
		t.Fatal("expected Next() to return true for single-node tree")
	}
	if iter.Current.Value != 42 {
		t.Errorf("expected Current.Value == 42, got %d", iter.Current.Value)
	}
	if iter.Next() {
		t.Errorf("expected Next() to return false after the only element")
	}
}

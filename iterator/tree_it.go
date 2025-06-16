package iterator

// Node represents a binary tree node with value, left/right children, and parent pointer
type Node struct {
	Value               int
	left, right, parent *Node
}

// NewNode creates a node with children, updating parent pointers
func NewNode(value int, left *Node, right *Node) (node *Node) {
	node = &Node{Value: value, left: left, right: right}
	if left != nil {
		left.parent = node
	}
	if right != nil {
		right.parent = node
	}
	return
}

// NewTerminalNode creates a node with no children
func NewTerminalNode(value int) *Node {
	return &Node{Value: value}
}

// InOrderIterator provides an iterator for in-order traversal
type InOrderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

// NewInOrderIterator creates and positions the iterator at the leftmost node
func NewInOrderIterator(root *Node) (inOrder *InOrderIterator) {
	inOrder = &InOrderIterator{
		Current:       root,
		root:          root,
		returnedStart: false,
	}
	if inOrder.Current != nil {
		for inOrder.Current.left != nil {
			inOrder.Current = inOrder.Current.left
		}
	}
	return
}

func (i *InOrderIterator) Reset() {
	i.Current = i.root
	i.returnedStart = false
	if i.Current != nil {
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
	}
}

// Next advances the iterator, returns true if there is a next element
func (i *InOrderIterator) Next() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true // first element
	}

	// If right child exists, go there and then all the way left
	if i.Current.right != nil {
		i.Current = i.Current.right
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
		return true
	} else {
		// Walk up to parent until we come from left
		p := i.Current.parent
		for p != nil && i.Current == p.right {
			i.Current = p
			p = p.parent
		}
		i.Current = p
		return i.Current != nil
	}
}

// BinaryTree holds the root node
type BinaryTree struct {
	root *Node
}

// NewBinaryTree creates a tree from a root node
func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}

// InOrder returns a new in-order iterator
func (b *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(b.root)
}

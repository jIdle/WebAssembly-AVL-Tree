package avl

import "math"

// setBalance determines and sets the balance factor of the node receiver.
// Balance factor is computed from the heights of the receiver's children.
func (n *node) setBalance() {
	if n.left == nil && n.right == nil {
		n.balance = 0
	} else if n.left == nil {
		n.balance = -n.right.height
	} else if n.right == nil {
		n.balance = n.left.height
	} else {
		n.balance = n.left.height - n.right.height
	}
}

// setHeight determines and sets the height of the node receiver.
// Height is computed by adding 1 to the larger height of the left and right child.
func (n *node) setHeight() {
	if n.left == nil && n.right == nil {
		n.height = 1
	} else if n.left == nil {
		n.height = n.right.height + 1
	} else if n.right == nil {
		n.height = n.left.height + 1
	} else {
		n.height = math.Max(n.left.height, n.right.height) + 1
	}
}

// GetSize returns the number of nodes in the tree.
func (t *AVL) GetSize() int {
	return t.size
}

// GetHeight returns the height of the tree.
func (t *AVL) GetHeight() float64 {
	return t.root.height
}

// rotateLeft will perform a left rotation given a parent and child node.
// The new parent node is returned.
func (t *AVL) rotateLeft(parent *node, child *node) *node {
	parent.right = child.left
	child.left = parent

	parent.setHeight()
	parent.setBalance()
	child.setHeight()
	child.setBalance()

	return child
}

// rotateRight will perform a right rotation given a parent and child node.
// The new parent node is returned.
func (t *AVL) rotateRight(parent *node, child *node) *node {
	parent.left = child.right
	child.right = parent

	parent.setHeight()
	parent.setBalance()
	child.setHeight()
	child.setBalance()

	return child
}

// checkBalance performs a rotation if the height-balance property is not satisfied.
// This is accomplished by correctly setting the current node's height and balance.
// Then depending on its balance and that of its children, a particular rotation
// will be performed, restoring the height-balance property of the AVL tree.
func (t *AVL) checkBalance(root *node) *node {
	root.setHeight()
	root.setBalance()

	if root.balance < -1 { // right rotation
		if root.right.balance > 0 {
			root.right = t.rotateRight(root.right, root.right.left) // double rotation
		}
		root = t.rotateLeft(root, root.right)
	} else if root.balance > 1 { // left rotation
		if root.left.balance < 0 {
			root.left = t.rotateLeft(root.left, root.left.right) // double rotation
		}
		root = t.rotateRight(root, root.left)
	}

	return root
}

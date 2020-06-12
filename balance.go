package main

import "math"

// setBalance : Called by node type. The node will ask its children for height values to compute its own balance factor.
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

// Height : Public function returns the total height of the tree. Calls recursive height function on root.
func (t *AVL) Height() float64 {
	if t.root == nil {
		return 0
	}
	return t.height(t.root)
}

// height : Private recursive function, returns height of specified node.
func (t *AVL) height(root *node) float64 {
	if root == nil {
		return 0
	}
	return math.Max(t.height(root.left), t.height(root.right)) + 1
}

// rotateLeft : Called by AVL type. Given a parent and child node, a left rotation is performed. The new parent node is returned.
func (t *AVL) rotateLeft(parent *node, child *node) *node {
	parent.right = child.left
	child.left = parent

	parent.setHeight()
	parent.setBalance()
	child.setHeight()
	child.setBalance()

	return child
}

// rotateRight : Called by AVL type. Given a parent and child node, a right rotation is performed. The new parent node is returned.
func (t *AVL) rotateRight(parent *node, child *node) *node {
	parent.left = child.right
	child.right = parent

	parent.setHeight()
	parent.setBalance()
	child.setHeight()
	child.setBalance()

	return child
}

// checkBalance : Called by AVL type. Correctly sets the given node's height and balance, then rotates if necessary. New parent node is returned.
func (t *AVL) checkBalance(root *node) *node {
	root.setHeight()
	root.setBalance()
	//fmt.Printf("Node: %v | Height: %v | Balance: %v\n", root.data, root.height, root.balance)

	if root.balance < -1 { // right rotation
		if root.right.balance > 0 {
			root.right = t.rotateRight(root.right, root.right.left)
		}
		root = t.rotateLeft(root, root.right)
	} else if root.balance > 1 { // left rotation
		if root.left.balance < 0 {
			root.left = t.rotateLeft(root.left, root.left.right)
		}
		root = t.rotateRight(root, root.left)
	}

	return root
}

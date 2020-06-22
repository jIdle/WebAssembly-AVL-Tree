package avl

// Remove is a wrapper function for the recursive remove function.
// It converts the argument data (basicData) passed in into a
// generic form, then passes it to the recursive remove function.
func (t *AVL) Remove(basicData interface{}) bool {
	data := checkType(basicData)
	if t.root == nil {
		return false
	}
	var success bool
	t.root, success = t.remove(t.root, data)
	return success
}

// remove is a recursive function which removes a node containing
// the user-specified data from the AVL tree. Location of removal
// is determined via binary search. A node pointer is returned to
// ensure that the changes propagate to the root node and persist.
// A bool is returned to signify whether the removal was a success
// or failure.
func (t *AVL) remove(root *node, data Interface) (*node, bool) {
	if root == nil {
		return nil, true
	}

	var success bool
	if data.Less(root.data) {
		root.left, success = t.remove(root.left, data)
	} else if root.data.Less(data) {
		root.right, success = t.remove(root.right, data)
	} else {
		t.size--
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			var ios *node
			root.right, ios = t.findIOS(root.right)
			ios.left = root.left
			ios.right = root.right
			root = ios
		}
		return root, true
	}

	return t.checkBalance(root), success
}

// findIOS is a helper function used by remove to search and return
// the In-Order Successor as replacement for the node marked for deletion.
func (t *AVL) findIOS(root *node) (*node, *node) {
	var ios *node
	if root.left == nil {
		ios := root
		root = root.right
		return root, ios
	}
	root.left, ios = t.findIOS(root.left)
	return t.checkBalance(root), ios
}

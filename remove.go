package main

// Remove : Wrapper function for Tree recursive remove.
func (t *AVL) Remove(basicData interface{}) bool {
	data := checkType(basicData)
	if t.root == nil {
		return false
	}
	var success bool
	t.root, success = t.remove(t.root, data)
	return success
}

// remove : Called by Tree type. Recursive binary removal. Returns error if applicable.
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

// findIOS : Helper function for remove to search and return the In-Order Successor
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

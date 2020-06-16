package avl

// Insert is a wrapper function for the recursive insert function.
// It converts the argument data (basicData) passed in into a
// generic form, then passes it to the recursive insert function.
func (t *AVL) Insert(basicData interface{}) {
	data := checkType(basicData)
	if t.root == nil {
		t.size++
		t.root = &node{data: data, height: 1, balance: 0, left: nil, right: nil}
		return
	}
	t.root = t.insert(t.root, data)
}

// insert is a recursive function which inserts a node containing
// the user-specified data into the AVL tree. Location of insertion
// is determined by binary search. A node pointer is returned to ensure
// that the changes propagate to the root node and persist.
func (t *AVL) insert(root *node, data Interface) *node {
	if root == nil {
		t.size++
		return &node{data: data, height: 1, balance: 0, left: nil, right: nil}
	} else if data.Less(root.data) {
		root.left = t.insert(root.left, data)
	} else if !data.Less(root.data) {
		root.right = t.insert(root.right, data)
	}

	return t.checkBalance(root)
}

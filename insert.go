package main

// Insert : Wrapper function for node insert.
func (t *Tree) Insert(basicData interface{}) {
	data := checkType(basicData)
	if t.root == nil {
		t.root = &node{data: data, height: 1, balance: 0, left: nil, right: nil}
		return
	}
	t.root = t.insert(t.root, data)
}

// insert : Called by Tree type. Recursive binary insertion. No return, should always insert.
func (t *Tree) insert(root *node, data Interface) *node {
	if root == nil {
		return &node{data: data, height: 1, balance: 0, left: nil, right: nil}
	} else if data.Less(root.data) {
		root.left = t.insert(root.left, data)
	} else if !data.Less(root.data) {
		root.right = t.insert(root.right, data)
	}

	return t.checkBalance(root)
}

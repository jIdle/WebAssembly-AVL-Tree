package avl

// Retrieve is a wrapper function for recursive retrieve function.
// It converts the argument data (basicData) passed in into a
// generic form, then passes it to the recursive retrieve function.
func (t *AVL) Retrieve(basicData interface{}) interface{} {
	data := checkType(basicData)
	if t.root == nil {
		return nil
	}
	return t.retrieve(t.root, data)
}

// retrieve is a recursive function which finds a node containing
// matching data to that user-specified data which was passed in
// via argument. Location of matching data is determined via binary
// search. An Interface type is returned since that is the form
// in which data is stored within the AVL tree. Nil is returned
// in the event that no matching data is discovered.
func (t *AVL) retrieve(root *node, data Interface) Interface {
	if root == nil {
		return nil
	}

	if data.Less(root.data) {
		return t.retrieve(root.left, data)
	} else if root.data.Less(data) {
		return t.retrieve(root.right, data)
	} else {
		return root.data
	}
}

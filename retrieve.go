package main

// Retrieve : Wrapper function for AVL recursive search.
func (t *AVL) Retrieve(basicData interface{}) interface{} {
	data := checkType(basicData)
	if t.root == nil {
		return nil
	}
	return t.retrieve(t.root, data)
}

// retrieve : Called by AVL type. Recursive binary search. Returns matching data.
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

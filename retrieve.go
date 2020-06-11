package main

import "errors"

// Retrieve : Wrapper function for Tree recursive search.
func (t *Tree) Retrieve(basicData interface{}) (Interface, error) {
	data := checkType(basicData)
	if t.root == nil {
		return nil, errors.New("no data to search in empty tree")
	}
	return t.retrieve(t.root, data)
}

// retrieve : Called by Tree type. Recursive binary search. Returns matching data.
func (t *Tree) retrieve(root *node, data Interface) (Interface, error) {
	if root == nil {
		return nil, errors.New("could not find specified value")
	}

	if data.Less(root.data) {
		return t.retrieve(root.left, data)
	} else if root.data.Less(data) {
		return t.retrieve(root.right, data)
	} else {
		return root.data, nil
	}
}

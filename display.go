package avl

import (
	"container/list"
)

// Ascending is a wrapper function for the recursive ascending traversal function.
func (t *AVL) Ascending() []interface{} {
	if t.root == nil {
		return nil
	}
	var order []interface{}
	return t.ascending(t.root, order)
}

// ascending is a recursive function which traverses the AVL tree
// in-order, appending each element to a slice, and returning that
// slice once finished.
func (t *AVL) ascending(root *node, order []interface{}) []interface{} {
	if root == nil {
		return order
	}
	return t.ascending(root.right, append(t.ascending(root.left, order), root.data))
}

// Descending is a wrapper function for recursive descending traversal function.
func (t *AVL) Descending() []interface{} {
	if t.root == nil {
		return nil
	}
	var order []interface{}
	return t.descending(t.root, order)
}

// descending is a recursive function which traverses the AVL tree
// reverse in-order, appending each element to a slice, and returning
// that slice once finished.
func (t *AVL) descending(root *node, order []interface{}) []interface{} {
	if root == nil {
		return order
	}
	return t.descending(root.left, append(t.descending(root.right, order), root.data))
}

// Preorder is a wrapper function for recursive preorder traversal function.
func (t *AVL) Preorder() []interface{} {
	if t.root == nil {
		return nil
	}
	var order []interface{}
	return t.preorder(t.root, order)
}

// preorder is a recursive function which traverses the AVL tree
// pre-order, appending each element to a slice, and returning that
// slice once finished.
func (t *AVL) preorder(root *node, order []interface{}) []interface{} {
	if root == nil {
		return order
	}
	return t.preorder(root.right, t.preorder(root.left, append(order, root.data)))
}

// Postorder is a wrapper function for recursive postorder traversal function.
func (t *AVL) Postorder() []interface{} {
	if t.root == nil {
		return nil
	}
	var order []interface{}
	return t.postorder(t.root, order)
}

// postorder is a recursive function which traverses the AVL tree
// post-order, appending each element to a slice, and returning that
// slice once finished.
func (t *AVL) postorder(root *node, order []interface{}) []interface{} {
	if root == nil {
		return order
	}
	return append(t.postorder(root.right, t.postorder(root.left, order)), root.data)
}

// LevelOrder is an iterative function which traverses the AVL tree
// in level-order, appending each element within a level to a slice,
// and appending each slice to another slice. This slice of slices is
// returned once the traversal is finished.
// In normal terms, the return type is just a 2D-array where each array
// represents a level, and the elements within each array represent
// the elements within a level.
// This traversal is achieved by using the BFS algorithm.
func (t *AVL) LevelOrder() []interface{} {
	if t.root == nil {
		return nil
	}

	type pair struct {
		first, second interface{}
	}

	var order []interface{}
	var o []interface{}

	lastLevel := 0
	queue := list.New()
	queue.PushBack(pair{t.root, 0})

	for queue.Len() != 0 {
		nlPair, _ := (queue.Remove(queue.Front())).(pair) // nl = node-level
		root, _ := nlPair.first.(*node)
		level, _ := nlPair.second.(int)

		if lastLevel != level {
			lastLevel = level
			order = append(order, o)
			o = nil
		}
		o = append(o, root.data)

		if root.left != nil {
			queue.PushBack(pair{root.left, level + 1})
		}
		if root.right != nil {
			queue.PushBack(pair{root.right, level + 1})
		}
	}
	order = append(order, o)

	return order
}

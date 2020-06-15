package main

import (
	"container/list"
)

// Ascending : Wrapper function for recursive ascending display
func (t *AVL) Ascending() []interface{} {
	if t.root == nil {
		return nil
	}
	var order []interface{}
	return t.ascending(t.root, order)
}

// ascending : Called by tree type. Recursive ascending display. Returns number of nodes.
func (t *AVL) ascending(root *node, order []interface{}) []interface{} {
	if root == nil {
		return order
	}
	return t.ascending(root.right, append(t.ascending(root.left, order), root.data))
}

// Descending : Wrapper function for recursive descending display.
func (t *AVL) Descending() []interface{} {
	if t.root == nil {
		return nil
	}
	var order []interface{}
	return t.descending(t.root, order)
}

// descending : Called by tree type. Recursive descending display. Returns number of nodes.
func (t *AVL) descending(root *node, order []interface{}) []interface{} {
	if root == nil {
		return order
	}
	return t.descending(root.left, append(t.descending(root.right, order), root.data))
}

// Preorder : Wrapper function for recursive preorder display.
func (t *AVL) Preorder() []interface{} {
	if t.root == nil {
		return nil
	}
	var order []interface{}
	return t.preorder(t.root, order)
}

// preorder : Called by tree type. Recursive preorder display. Returns number of nodes.
func (t *AVL) preorder(root *node, order []interface{}) []interface{} {
	if root == nil {
		return order
	}
	return t.preorder(root.right, t.preorder(root.left, append(order, root.data)))
}

// LevelOrder : Uses the BFS algorithm to display nodes in level order. Returns number of nodes.
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

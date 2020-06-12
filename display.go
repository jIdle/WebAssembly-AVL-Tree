package main

import (
	"container/list"
	"fmt"
)

// Ascending : Wrapper function for recursive ascending display
func (t *AVL) Ascending() int {
	if t.root == nil {
		return 0
	}
	return t.ascending(t.root)
}

// ascending : Called by tree type. Recursive ascending display. Returns number of nodes.
func (t *AVL) ascending(root *node) int {
	if root == nil {
		return 0
	}
	retVal := t.ascending(root.left) + 1
	fmt.Printf("%v ", root.data)
	return t.ascending(root.right) + retVal
}

// Descending : Wrapper function for recursive descending display.
func (t *AVL) Descending() int {
	if t.root == nil {
		return 0
	}
	return t.descending(t.root)
}

// descending : Called by tree type. Recursive descending display. Returns number of nodes.
func (t *AVL) descending(root *node) int {
	if root == nil {
		return 0
	}
	retVal := t.descending(root.right) + 1
	fmt.Printf("%v ", root.data)
	return t.descending(root.left) + retVal
}

// Preorder : Wrapper function for recursive preorder display.
func (t *AVL) Preorder() int {
	if t.root == nil {
		return 0
	}
	return t.preorder(t.root)
}

// preorder : Called by tree type. Recursive preorder display. Returns number of nodes.
func (t *AVL) preorder(root *node) int {
	if root == nil {
		return 0
	}
	fmt.Printf("%v ", root.data)
	return t.preorder(root.left) + t.preorder(root.right) + 1
}

// LevelOrder : Uses the BFS algorithm to display nodes in level order. Returns number of nodes.
func (t *AVL) LevelOrder(showBalance bool) int {
	if t.root == nil {
		return 0
	}

	type pair struct {
		first, second interface{}
	}

	count := 0
	lastLevel := 1
	queue := list.New()
	queue.PushBack(pair{t.root, 1})

	fmt.Printf("Level %v: ", 1)
	for queue.Len() != 0 {
		count++
		nlPair, _ := (queue.Remove(queue.Front())).(pair) // nl = node-level
		root, _ := nlPair.first.(*node)
		level, _ := nlPair.second.(int)

		if lastLevel != level {
			lastLevel = level
			fmt.Printf("\nLevel %v: ", level)
		}
		if showBalance {
			//fmt.Printf("%v|%v ", js.Value(root.data.(Object)).Get("key").Int(), root.balance)
			fmt.Printf("%v|%v ", root.data, root.balance)
		} else {
			//fmt.Printf("%v ", js.Value(root.data.(Object)).Get("key").Int())
			fmt.Printf("%v ", root.data)
		}

		if root.left != nil {
			queue.PushBack(pair{root.left, level + 1})
		}
		if root.right != nil {
			queue.PushBack(pair{root.right, level + 1})
		}
	}
	fmt.Println()

	return count
}

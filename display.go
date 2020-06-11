package main

import (
	"container/list"
	"fmt"
)

// AscendingDisplay : Wrapper function for recursive ascending display
func (t *Tree) AscendingDisplay() int {
	if t.root == nil {
		return 0
	}
	return t.ascending(t.root)
}

// ascending : Called by tree type. Recursive ascending display. Returns number of nodes.
func (t *Tree) ascending(root *node) int {
	if root == nil {
		return 0
	}
	retVal := t.ascending(root.left) + 1
	fmt.Printf("%d ", root.data)
	return t.ascending(root.right) + retVal
}

// DescendingDisplay : Wrapper function for recursive descending display.
func (t *Tree) DescendingDisplay() int {
	if t.root == nil {
		return 0
	}
	return t.descending(t.root)
}

// descending : Called by tree type. Recursive descending display. Returns number of nodes.
func (t *Tree) descending(root *node) int {
	if root == nil {
		return 0
	}
	retVal := t.descending(root.right) + 1
	fmt.Printf("%d ", root.data)
	return t.descending(root.left) + retVal
}

// PreorderDisplay : Wrapper function for recursive preorder display.
func (t *Tree) PreorderDisplay() int {
	if t.root == nil {
		return 0
	}
	return t.preorder(t.root)
}

// preorder : Called by tree type. Recursive preorder display. Returns number of nodes.
func (t *Tree) preorder(root *node) int {
	if root == nil {
		return 0
	}
	fmt.Printf("%d ", root.data)
	return t.preorder(root.left) + t.preorder(root.right) + 1
}

// LevelDisplay : Uses the BFS algorithm to display nodes in level order. Returns number of nodes.
func (t *Tree) LevelDisplay(showBalance bool) int {
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
			fmt.Printf("%d|%v ", root.data, root.balance)
		} else {
			fmt.Printf("%d ", root.data)
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

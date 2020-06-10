package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
)

var scanner = bufio.NewScanner(os.Stdin)

// Interface : Allows comparison between keys in generic form.
type Interface interface {
	Less(toCompare Interface) bool
}

type Int int
type Float float64
type String string

// Less : Assume a, b are Int. Less returns true if a < b.
func (a Int) Less(b Interface) bool {
	return a < b.(Int)
}

// Less : Assume a, b are Float64. Less returns true if a < b.
func (a Float) Less(b Interface) bool {
	return a < b.(Float)
}

// Less : Assume a, b are String. Less returns true if a < b.
func (a String) Less(b Interface) bool {
	return a < b.(String)
}

func checkType(toConvert interface{}) Interface {
	switch t := toConvert.(type) {
	case int:
		return Int(t)
	case Int:
		return t
	case float64:
		return Float(t)
	case Float:
		return t
	case string:
		return String(t)
	case String:
		return t
	default:
		panic(fmt.Sprintf("Could not find a compatible type conversion for %v\n", reflect.TypeOf(t).Name()))
	}
}

// node : Container for user data
type node struct { // Should not be exported. Only tree data structure should have knowledge of a node.
	data    Interface
	height  float64
	balance float64
	left    *node
	right   *node
}

// setBalance : Called by node type. The node will ask its children for height values to compute its own balance factor.
func (n *node) setBalance() {
	if n.left == nil && n.right == nil {
		n.balance = 0
	} else if n.left == nil {
		n.balance = -n.right.height
	} else if n.right == nil {
		n.balance = n.left.height
	} else {
		n.balance = n.left.height - n.right.height
	}
}

func (n *node) setHeight() {
	if n.left == nil && n.right == nil {
		n.height = 1
	} else if n.left == nil {
		n.height = n.right.height + 1
	} else if n.right == nil {
		n.height = n.left.height + 1
	} else {
		n.height = math.Max(n.left.height, n.right.height) + 1
	}
}

// Tree : Container for root node.
type Tree struct {
	root *node
}

// Height : Public function returns the total height of the tree. Calls recursive height function on root.
func (t *Tree) Height() float64 {
	if t.root == nil {
		return 0
	}
	return t.height(t.root)
}

// height : Private recursive function, returns height of specified node.
func (t *Tree) height(root *node) float64 {
	if root == nil {
		return 0
	}
	return math.Max(t.height(root.left), t.height(root.right)) + 1
}

// rotateLeft : Called by Tree type. Given a parent and child node, a left rotation is performed. The new parent node is returned.
func (t *Tree) rotateLeft(parent *node, child *node) *node {
	parent.right = child.left
	child.left = parent

	parent.setHeight()
	parent.setBalance()
	child.setHeight()
	child.setBalance()

	return child
}

// rotateRight : Called by Tree type. Given a parent and child node, a right rotation is performed. The new parent node is returned.
func (t *Tree) rotateRight(parent *node, child *node) *node {
	parent.left = child.right
	child.right = parent

	parent.setHeight()
	parent.setBalance()
	child.setHeight()
	child.setBalance()

	return child
}

// checkBalance : Called by Tree type. Correctly sets the given node's height and balance, then rotates if necessary. New parent node is returned.
func (t *Tree) checkBalance(root *node) *node {
	root.setHeight()
	root.setBalance()
	//fmt.Printf("Node: %v | Height: %v | Balance: %v\n", root.data, root.height, root.balance)

	if root.balance < -1 { // right rotation
		if root.right.balance > 0 {
			root.right = t.rotateRight(root.right, root.right.left)
		}
		root = t.rotateLeft(root, root.right)
	} else if root.balance > 1 { // left rotation
		if root.left.balance < 0 {
			root.left = t.rotateLeft(root.left, root.left.right)
		}
		root = t.rotateRight(root, root.left)
	}

	return root
}

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

// Remove : Wrapper function for Tree recursive remove.
func (t *Tree) Remove(basicData interface{}) error {
	data := checkType(basicData)
	if t.root == nil {
		return errors.New("no data to remove in empty tree")
	}
	var err error
	t.root, err = t.remove(t.root, data)
	return err
}

// remove : Called by Tree type. Recursive binary removal. Returns error if applicable.
func (t *Tree) remove(root *node, data Interface) (*node, error) {
	var err error
	if root == nil {
		err = errors.New("could not find specified value")
		return nil, err
	}

	if data.Less(root.data) {
		root.left, err = t.remove(root.left, data)
	} else if root.data.Less(data) {
		root.right, err = t.remove(root.right, data)
	} else {
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
		return root, nil
	}

	return t.checkBalance(root), err
}

// findIOS : Helper function for remove to search and return the In-Order Successor
func (t *Tree) findIOS(root *node) (*node, *node) {
	var ios *node
	if root.left == nil {
		ios := root
		root = root.right
		return root, ios
	}
	root.left, ios = t.findIOS(root.left)
	return t.checkBalance(root), ios
}

// Search : Wrapper function for Tree recursive search.
func (t *Tree) Search(basicData interface{}) (Interface, error) {
	data := checkType(basicData)
	if t.root == nil {
		return nil, errors.New("no data to search in empty tree")
	}
	return t.search(t.root, data)
}

// search : Called by Tree type. Recursive binary search. Returns matching data.
func (t *Tree) search(root *node, data Interface) (Interface, error) {
	if root == nil {
		return nil, errors.New("could not find specified value")
	}

	if data.Less(root.data) {
		return t.search(root.left, data)
	} else if root.data.Less(data) {
		return t.search(root.right, data)
	} else {
		return root.data, nil
	}
}

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

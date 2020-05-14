// WAVL Tree Implementation
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

type node struct { // Should not be exported. Only BST data structure should have knowledge of a node.
	data  int
	left  *node
	right *node
}

// insert : Called by node type. Integer argument will be inserted in tree. Error is return if root node is nil.
func (cur *node) insert(data int) error {
	if cur == nil {
		return errors.New("initialize root node first")
	}

	if data < cur.data {
		if cur.left == nil {
			cur.left = &node{data: data}
			return nil
		}
		return cur.left.insert(data)
	} else if data >= cur.data {
		if cur.right == nil {
			cur.right = &node{data: data}
			return nil
		}
		return cur.right.insert(data)
	}
	return nil
}

// findIOS : Helper Function : Called by node type. Called by right child of the node to be removed. Goes left until nil from there. removes and returns farthest left node as IOS to node-to-be-removed.
//func (cur *node) findIOS(par *node) int {
//	if cur.left == nil {
//		par.left = cur.right
//		return cur.data
//	}
//	return cur.left.findIOS(cur)
//}

// replaceChild : Helper function for remove. This function will remove a node from a bst by replacing it with repl.
//func (cur *node) replaceChild(par *node, repl *node) error {
//	if cur == nil {
//		return errors.New("The function replaceChild() can't be called on a nil node.")
//	}
//	if par == nil {
//		return errors.New("Removal of root node case not caught in wrapper function.")
//	}
//
//	if par.left == cur {
//		par.left = repl
//	} else {
//		par.right = repl
//	}
//	return nil
//}

// remove : Wrapper function for _remove.
//func (cur *node) remove(data int) error {
//	return cur._remove(data, nil)
//}

// _remove : Called by node type. Integer argument will be removed if found. Otherwise return error.
//func (cur *node) _remove(data int, par *node) error {
//	if cur == nil {
//		return errors.New("Could not find specified value.")
//	}
//
//	switch {
//	case data < cur.data:
//		return cur.left._remove(data, cur)
//	case data > cur.data:
//		return cur.right._remove(data, cur)
//	case data == cur.data:
//		if cur.left == nil && cur.right == nil {
//			return cur.replaceChild(par, nil)
//		} else if cur.left == nil {
//			return cur.replaceChild(par, cur.right)
//		} else if cur.right == nil {
//			return cur.replaceChild(par, cur.left)
//		}
//		// Check that right has a left to warrant calling IOS
//		rChild := cur.right
//		if rChild.left == nil {
//			cur.data = rChild.data
//			cur.right = rChild.right
//		} else {
//			cur.data = rChild.findIOS(nil)
//		}
//	}
//	return nil
//}

// Tree : Container for root node.
type Tree struct {
	root *node
}

// Insert : Wrapper function for node insert.
func (t *Tree) Insert(data int) error {
	if t.root == nil {
		t.root = &node{data: data}
		return nil
	}
	return t.root.insert(data)
}

// Find : Wrapper function for node find.
func (t *Tree) Find(data int) (int, error) {
	if t.root == nil {
		return 0, errors.New("no data to find in empty tree")
	}
	return t.find(t.root, data)
}

func (t *Tree) find(current *node, data int) (int, error) {
	if current == nil {
		return 0, errors.New("could not find specified value")
	}
	if data == current.data {
		return current.data, nil
	} else if data < current.data {
		return t.find(current.left, data)
	}
	return t.find(current.right, data)
}

// Remove : ...
func (t *Tree) Remove(data int) error {
	if t.root == nil {
		return errors.New("No data to remove in empty tree.")
	}

}

// Remove : Wrapper function for node remove.
//func (t *Tree) Remove(data int) error {
//	if t.root == nil {
//		return errors.New("No data to remove in empty tree.")
//	}
//
//	if data != t.root.data {
//		return t.root.remove(data)
//	}
//	switch {
//	case t.root.left == nil && t.root.right == nil:
//		t.root = nil
//	case t.root.left == nil:
//		t.root = t.root.right
//	case t.root.right == nil:
//		t.root = t.root.left
//	default:
//		return t.root.remove(data)
//	}
//	return nil
//}

// Display : Wrapper function for recursive display()
func (t *Tree) Display() int {
	if t.root == nil {
		return 0
	}
	return t.display(t.root, 0)
}

// display : Called by tree type. Recursive preorder display. Returns number of nodes.
func (t *Tree) display(current *node, level int) int {
	if current == nil {
		return 0
	}
	retVal := t.display(current.left, level+1) + 1
	fmt.Printf("Level %d: %d\n", level, current.data) // probably stupid and will remove
	return t.display(current.right, level+1) + retVal
}

func main() {
	myTree := &Tree{}
	for {
		fmt.Println("\n\t0) Exit")
		fmt.Println("\t1) Insert")
		fmt.Println("\t2) Find")
		fmt.Println("\t3) Display")
		fmt.Println("\t4) Remove")
		scanner.Scan()
		fmt.Println()
		switch input, _ := strconv.Atoi(scanner.Text()); input {
		case 0:
			fmt.Println("Exiting...")
			os.Exit(1)
		case 1:
			fmt.Println("Enter the value you would like to insert:")
			scanner.Scan()
			input, _ = strconv.Atoi(scanner.Text())
			myTree.Insert(input)
		case 2:
			fmt.Println("Enter the number you would like to search for:")
			scanner.Scan()
			input, _ = strconv.Atoi(scanner.Text())
			if data, e := myTree.Find(input); e == nil {
				fmt.Printf("%d was found in the tree.\n", data)
			} else {
				fmt.Println(e)
			}
		case 3:
			fmt.Printf("There are %d node(s) in the tree.\n", myTree.Display())
		case 4:
			fmt.Println("Enter the value you would like to remove:")
			scanner.Scan()
			input, _ = strconv.Atoi(scanner.Text())
			if e := myTree.Remove(input); e != nil {
				fmt.Println(e)
			} else {
				fmt.Println("Value successfully removed from tree.")
			}
		default:
			fmt.Println("Please enter a valid input.")
		}
	}
}

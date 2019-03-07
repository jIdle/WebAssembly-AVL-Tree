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
	Data  int
	Left  *node
	Right *node
}

// insert : Called by node type. Integer argument will be inserted in tree. Error is return if root node is nil.
func (cur *node) insert(data int) error {
	if cur == nil {
		return errors.New("Initialize root node first.")
	}

	if data < cur.Data {
		if cur.Left == nil {
			cur.Left = &node{Data: data}
			return nil
		}
		return cur.Left.insert(data)
	} else if data >= cur.Data {
		if cur.Right == nil {
			cur.Right = &node{Data: data}
			return nil
		}
		return cur.Right.insert(data)
	}
	return nil
}

// findIOS : Helper Function : Called by node type. Called by right child of the node to be removed. Goes left until nil from there. removes and returns farthest left node as IOS to node-to-be-removed.
func (cur *node) findIOS(par *node) int {
	if cur.Left == nil {
		par.Left = cur.Right
		return cur.Data
	}
	return cur.Left.findIOS(cur)
}

// replaceChild : Helper function for remove. This function will remove a node from a bst by replacing it with repl.
func (cur *node) replaceChild(par *node, repl *node) error {
	if cur == nil {
		return errors.New("The function replaceChild() can't be called on a nil node.")
	}
	if par == nil {
		return errors.New("Removal of root node case not caught in wrapper function.")
	}

	if par.Left == cur {
		par.Left = repl
	} else {
		par.Right = repl
	}
	return nil
}

// remove : Wrapper function for _remove.
func (cur *node) remove(data int) error {
	return cur._remove(data, nil)
}

// _remove : Called by node type. Integer argument will be removed if found. Otherwise return error.
func (cur *node) _remove(data int, par *node) error {
	if cur == nil {
		return errors.New("Could not find specified value.")
	}

	switch {
	case data < cur.Data:
		return cur.Left._remove(data, cur)
	case data > cur.Data:
		return cur.Right._remove(data, cur)
	case data == cur.Data:
		if cur.Left == nil && cur.Right == nil {
			return cur.replaceChild(par, nil)
		} else if cur.Left == nil {
			return cur.replaceChild(par, cur.Right)
		} else if cur.Right == nil {
			return cur.replaceChild(par, cur.Left)
		}
		// Check that right has a left to warrant calling IOS
		rChild := cur.Right
		if rChild.Left == nil {
			cur.Data = rChild.Data
			cur.Right = rChild.Right
		} else {
			cur.Data = rChild.findIOS(nil)
		}
	}
	return nil
}

// find : Called by node type. Integer argument will be returned search for in tree. Matching integer in tree will be returned with success boolean, if found. Otherwise 0 and false are returned.
func (cur *node) find(data int) (int, error) {
	if cur == nil {
		return 0, errors.New("Could not find specified value.")
	}
	if data == cur.Data {
		return cur.Data, nil
	} else if data < cur.Data {
		return cur.Left.find(data)
	} else {
		return cur.Right.find(data)
	}
}

// displayAll : Wrapper function for recursive _displayAll().
func (cur *node) displayAll() int {
	return cur._displayAll(0)
}

// _displayAll : Called by node type. Recursive version of display. No arguments taken, and an integer is returned to represent the number of nodes displayed (total # of nodes).
func (cur *node) _displayAll(level int) int {
	if cur == nil {
		return 0
	}
	retVal := cur.Left._displayAll(level+1) + 1
	fmt.Printf("Level %d: %d\n", level, cur.Data)
	return cur.Right._displayAll(level+1) + retVal
}

// Tree : Container for root node.
type Tree struct {
	Root *node
}

// Insert : Wrapper function for node insert.
func (t *Tree) Insert(data int) error {
	if t.Root == nil {
		t.Root = &node{Data: data}
		return nil
	}
	return t.Root.insert(data)
}

// Find : Wrapper function for node find.
func (t *Tree) Find(data int) (int, error) {
	if t.Root == nil {
		return 0, errors.New("No data to find in empty tree.")
	}
	return t.Root.find(data)
}

// Remove : Wrapper function for node remove.
func (t *Tree) Remove(data int) error {
	if t.Root == nil {
		return errors.New("No data to remove in empty tree.")
	}

	if data != t.Root.Data {
		return t.Root.remove(data)
	}
	switch {
	case t.Root.Left == nil && t.Root.Right == nil:
		t.Root = nil
	case t.Root.Left == nil:
		t.Root = t.Root.Right
	case t.Root.Right == nil:
		t.Root = t.Root.Left
	default:
		return t.Root.remove(data)
	}
	return nil
}

// Display : Wrapper function for node displayAll.
func (t *Tree) Display() int {
	if t.Root == nil {
		return 0
	}
	return t.Root.displayAll()
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

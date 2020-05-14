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

// node : Container for user data
type node struct { // Should not be exported. Only BST data structure should have knowledge of a node.
	data  int
	left  *node
	right *node
}

// Tree : Container for root node.
type Tree struct {
	root *node
}

// Insert : Wrapper function for node insert.
func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = &node{data: data}
		return
	}
	t.root = t.insert(t.root, data)
}

// insert : Called by Tree type. Recursive binary insertion. No return, should always insert.
func (t *Tree) insert(current *node, data int) *node {
	if current == nil {
		return &node{data: data}
	} else if data < current.data {
		current.left = t.insert(current.left, data)
	} else if data >= current.data {
		current.right = t.insert(current.right, data)
	}
	return current
}

// Remove : Wrapper function for Tree recursive remove.
func (t *Tree) Remove(data int) error {
	if t.root == nil {
		return errors.New("no data to remove in empty tree")
	}
	return t.remove(t.root, data)
}

// remove : Called by Tree type. Recursive binary removal. Returns error if applicable.
func (t *Tree) remove(current *node, data int) error {
	if current == nil {
		return errors.New("could not find specified value")
	}

	if data == current.data {
		if current.left == nil && current.right == nil {
			current = nil
		} else if current.left == nil {
			current = current.right
		} else if current.right == nil {
			current = current.left
		}
		current = t.findIOS(current.right)
	} else if data < current.data {
		return t.remove(current.left, data)
	}
	return t.remove(current.right, data)
}

// findIOS : Helper function for remove to find and return the In-Order Successor
func (t *Tree) findIOS(current *node) *node {
	if current == nil {
		return nil
	} else if current.left == nil {
		temp := current
		current = current.right
		return temp
	}
	return t.findIOS(current.left)
}

// Find : Wrapper function for Tree recursive find.
func (t *Tree) Find(data int) (int, error) {
	if t.root == nil {
		return 0, errors.New("no data to find in empty tree")
	}
	return t.find(t.root, data)
}

// find : Called by Tree type. Recursive binary search. Returns matching data.
func (t *Tree) find(current *node, data int) (int, error) {
	// "data" in this case may not always be an integer, could be a whole object
	// but we match for the identifier/name
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

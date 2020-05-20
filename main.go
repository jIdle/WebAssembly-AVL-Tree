// WAVL Tree Implementation
package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

// node : Container for user data
type node struct { // Should not be exported. Only BST data structure should have knowledge of a node.
	data  int
	rank  int
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
func (t *Tree) insert(root *node, data int) *node {
	if root == nil {
		return &node{data: data}
	} else if data < root.data {
		root.left = t.insert(root.left, data)
	} else if data >= root.data {
		root.right = t.insert(root.right, data)
	}
	return root
}

// Remove : Wrapper function for Tree recursive remove.
func (t *Tree) Remove(data int) error {
	if t.root == nil {
		return errors.New("no data to remove in empty tree")
	}
	var err error
	t.root, err = t.remove(t.root, data)
	return err
}

// remove : Called by Tree type. Recursive binary removal. Returns error if applicable.
func (t *Tree) remove(root *node, data int) (*node, error) {
	var err error
	if root == nil {
		err = errors.New("could not find specified value")
		return nil, err
	}

	if data == root.data {
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
	} else if data < root.data {
		root.left, err = t.remove(root.left, data)
		return root, err
	}
	root.right, err = t.remove(root.right, data)
	return root, err
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
	return root, ios
}

// Search : Wrapper function for Tree recursive search.
func (t *Tree) Search(data int) (int, error) {
	if t.root == nil {
		return 0, errors.New("no data to search in empty tree")
	}
	return t.search(t.root, data)
}

// search : Called by Tree type. Recursive binary search. Returns matching data.
func (t *Tree) search(root *node, data int) (int, error) {
	// "data" in this case may not always be an integer, could be a whole object
	// but we match for the identifier/name
	if root == nil {
		return 0, errors.New("could not find specified value")
	}
	if data == root.data {
		return root.data, nil
	} else if data < root.data {
		return t.search(root.left, data)
	}
	return t.search(root.right, data)
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
func (t *Tree) LevelDisplay() int {
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
		fmt.Printf("%d ", root.data)

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

func main() {
	myTree := &Tree{}
	for {
		fmt.Println("\nSelect an option from the menu:")
		fmt.Println("| 0) Exit")
		fmt.Println("| 1) Insert")
		fmt.Println("| 2) Search")
		fmt.Println("| 3) Display")
		fmt.Println("| 4) Remove")
		fmt.Println("| 5) Height")
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

			if data, e := myTree.Search(input); e != nil {
				fmt.Println("Error encountered: ", e)
			} else {
				fmt.Printf("%d was found in the tree.\n", data)
			}
		case 3:
			fmt.Println("Select a display type:")
			fmt.Println("| 1) Ascending")
			fmt.Println("| 2) Descending")
			fmt.Println("| 3) Preorder")
			fmt.Println("| 4) Level Order")
			scanner.Scan()
			fmt.Println()

			switch input, _ := strconv.Atoi(scanner.Text()); input {
			case 1:
				fmt.Printf("There are %d nodes in the tree.\n", myTree.AscendingDisplay())
			case 2:
				fmt.Printf("There are %d nodes in the tree.\n", myTree.DescendingDisplay())
			case 3:
				fmt.Printf("There are %d nodes in the tree.\n", myTree.PreorderDisplay())
			case 4:
				fmt.Printf("There are %d nodes in the tree.\n", myTree.LevelDisplay())
			default:
				fmt.Println("Please enter a valid input.")
			}
		case 4:
			fmt.Println("Enter the value you would like to remove:")
			scanner.Scan()
			input, _ = strconv.Atoi(scanner.Text())

			if e := myTree.Remove(input); e != nil {
				fmt.Println("Error encountered: ", e)
			} else {
				fmt.Println("Value successfully removed from tree.")
			}
		case 5:
			fmt.Println("Height of tree is: ", myTree.Height())
		default:
			fmt.Println("Please enter a valid input.")
		}
	}
}

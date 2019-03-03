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

type node struct {
	Data  int
	Left  *node
	Right *node
}

func (current *node) Insert(data int) error {
	if current == nil {
		return errors.New("Initialize root node first.")
	}

	if data == current.Data { // Fix this, needs to push data right on equal
		return nil
	} else if data < current.Data {
		if current.Left == nil {
			current.Left = &node{Data: data}
			return nil
		}
		return current.Left.Insert(data)
	} else if data >= current.Data {
		if current.Right == nil {
			current.Right = &node{Data: data}
			return nil
		}
		return current.Right.Insert(data)
	}
	return nil
}

func (current *node) Find(data int) (int, bool) {
	if current == nil {
		return 0, false
	}
	if data == current.Data {
		return current.Data, true
	} else if data < current.Data {
		return current.Left.Find(data)
	} else {
		return current.Right.Find(data)
	}
}

func (current *node) Display() int {
	if current == nil {
		return 0
	}
	retVal := current.Left.Display() + 1
	fmt.Println("Data: " + strconv.Itoa(current.Data))
	return current.Right.Display() + retVal
}

func main() { // For right now main is just used to test the functions incrementally
	head := node{10, nil, nil}
	for {
		fmt.Println("\n3) Find")
		fmt.Println("2) Display")
		fmt.Println("1) Insert")
		fmt.Println("0) Exit")
		scanner.Scan()
		fmt.Println()
		input, _ := strconv.Atoi(scanner.Text())
		if input == 1 {
			fmt.Println("Enter the value you would like to insert:")
			scanner.Scan()
			input, _ = strconv.Atoi(scanner.Text())
			head.Insert(input)
		} else if input == 2 {
			fmt.Printf("There are %d node(s) in the tree\n", head.Display())
		} else if input == 3 {
			fmt.Println("Enter the number you would like to search for:")
			scanner.Scan()
			input, _ = strconv.Atoi(scanner.Text())
			_, inTree := head.Find(input)
			if inTree == true {
				fmt.Println("Value found in tree.")
			} else {
				fmt.Println("Value not found in tree.")
			}
		} else if input == 0 {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Please enter a valid input.")
		}
	}
	os.Exit(1)
}

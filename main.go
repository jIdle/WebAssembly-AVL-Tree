// WAVL Tree Implementation
package main

import (
	"fmt"
	"os"
	"strconv"
)

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

// WAVL Tree Implementation
package main

import (
    "fmt"
    "bufio"
    "strconv"
    "os"
    "errors"
)

var scanner = bufio.NewScanner(os.Stdin)

type Node struct {
    Data int
    Left * Node
    Right * Node
}

func (current * Node) Insert(data int) error {
    if current == nil {
        return errors.New("Initialize root node first.")
    }

    if data == current.Data { // Fix this, needs to push data right on equal
        return nil
    } else if data < current.Data {
        if current.Left == nil {
            current.Left = &Node{Data: data}
            return nil
        }
        return current.Left.Insert(data)
    } else if data > current.Data {
        if current.Right == nil {
            current.Right = &Node{Data: data}
            return nil
        }
        return current.Right.Insert(data)
    }
    return nil
}

func (current * Node) Display() int {
    if current == nil {
        return 0
    }
    retVal := current.Left.Display() + 1
    fmt.Println("Data: " + strconv.Itoa(current.Data))
    return current.Right.Display() + retVal
}

func main() { // For right now main is just used to test the functions incrementally
    head := Node{10, nil, nil}
    var input int = 1
    for {
        fmt.Println("\n2) Display")
        fmt.Println("1) Insert")
        fmt.Println("0) Exit")
        scanner.Scan()
        fmt.Println()
        input, _ = strconv.Atoi(scanner.Text())
        if input == 1 {
            fmt.Println("Enter the value you would like to insert:")
            scanner.Scan()
            input, _ = strconv.Atoi(scanner.Text())
            head.Insert(input)
        } else if input == 2 {
            fmt.Printf("There are %d node(s) in the tree\n", head.Display())
        } else if input == 0 {
            fmt.Println("Exiting...")
            os.Exit(0)
        } else {
            fmt.Println("Please enter a valid input.")
        }
    }
    fmt.Println("Error, outside main loop.")
    os.Exit(1)
}
// WAVL Tree Implementation

import (
    "fmt"
    "bufio"
    "strconv"
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

    if data == current.Data {
        return nil
    } else if data < current.Data {
        if current.Left == nil {
            current.Left = &Node{Data: data}
            return nil
        }
        return current.Left.Insert(data)
    } else if data > current.Data {
        if current.Right == nil {
            n.Right = &Node{Data: data}
            return nil
        }
        return n.Right.Insert(data)
    }
    return nil
}

func (current * Node) Display() int {
    if current == nil {
        return 0
    }
    retVal := current.Left.Display() + 1
    fmt.Println("Data: " + current.Data)
    return current.Right.Display() + retVal
}

func main() {
    Node head
    userInput := 1
    for userInput == 1 {
        fmt.Println("1) Insert")
        fmt.Println("2) Display")
        fmt.Println("Which would you like to do?")
        scanner.Scan()
        if scanner.Text() == "1" {
            fmt.Println("Enter the value you would like to insert:")
            scanner.Scan()
            head.Insert(strconv.Atoi(scanner.Text()))
        } else if scanner.Text() == "2" {
            fmt.Println("Data in the bst:")
            head.Display()
        }
    }
}
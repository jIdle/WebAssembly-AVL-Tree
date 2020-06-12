package main

import (
	"fmt"
	"reflect"
	"syscall/js"
)

// Interface : Allows comparison between keys in generic form.
type Interface interface {
	Less(toCompare Interface) bool
}

// AVL : Container for root node.
type AVL struct {
	root *node
}

// NewAVL : Returns a new and empty AVL tree. May initialize JS wrapper functions.
func NewAVL() *AVL {
	tree := &AVL{root: nil}
	js.Global().Set("Insert", js.FuncOf(tree.InsertJS))
	js.Global().Set("Remove", js.FuncOf(tree.RemoveJS))
	js.Global().Set("Retrieve", js.FuncOf(tree.RetrieveJS))
	js.Global().Set("Display", js.FuncOf(tree.DisplayJS))
}

// node : Container for user data
type node struct {
	data    Interface
	height  float64
	balance float64
	left    *node
	right   *node
}

type Int int
type Float float64
type String string
type Object js.Value

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

// Less : Assume a, b are Object. Less returns true if a < b.
func (a Object) Less(b Interface) bool {
	return js.Value(a).Call("Less", (js.Value(b.(Object)))).Bool()
}

func checkType(toConvert interface{}) Interface {
	switch goVal := toConvert.(type) {
	case int:
		return Int(goVal)
	case float64:
		return Float(goVal)
	case string:
		return String(goVal)
	case js.Value:
		switch jsType := goVal.Type().String(); jsType {
		case "number":
			Number := js.Global().Get("Number")
			if Number.Call("isInteger", goVal).Bool() {
				return Int(goVal.Int())
			} else {
				return Float(goVal.Float())
			}
		case "string":
			return String(goVal.String())
		case "object":
			return Object(goVal)
		default:
			panic(fmt.Sprintf("Input type can not be ordered.\n"))
		}
	default:
		if v, ok := goVal.(Interface); ok {
			return v
		}
		panic(fmt.Sprintf("Could not find a compatible type conversion for %v\n", reflect.TypeOf(goVal).Name()))
	}
}

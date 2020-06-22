//  Package avl provides an implementation of a generic AVL tree.
//	The data structure is compatible not only with built-in and
//	user-defined types, but also Javascript types.
// 	The methods offered by this package are: Insert, Remove, and Retrieve
//	Data can be returned in ordered array format: In-Order, Reverse In-Order, Pre-Order, Post-Order, and Level-Order

package avl

import (
	"fmt"
	"reflect"
	"syscall/js"
)

// Interface type allows comparison between keys in generic form.
type Interface interface {
	Less(toCompare Interface) bool
}

// AVL type represents a container for an AVL tree.
type AVL struct {
	root *node
	size int
}

// NewAVL returns a new and empty AVL tree. Additionally, the exported AVL
// functions will be registered as callable Javascript functions. This
// allows the AVL library to function as a WASM module.
func NewAVL() *AVL {
	tree := &AVL{root: nil, size: 0}
	js.Global().Set("Insert", js.FuncOf(tree.InsertJS))
	js.Global().Set("Remove", js.FuncOf(tree.RemoveJS))
	js.Global().Set("Retrieve", js.FuncOf(tree.RetrieveJS))
	js.Global().Set("Ascending", js.FuncOf(tree.AscendingJS))
	js.Global().Set("Descending", js.FuncOf(tree.DescendingJS))
	js.Global().Set("Preorder", js.FuncOf(tree.PreorderJS))
	js.Global().Set("Postorder", js.FuncOf(tree.PostorderJS))
	js.Global().Set("LevelOrder", js.FuncOf(tree.LevelOrderJS))
	return tree
}

// node is a container for user data.
type node struct {
	data    Interface
	height  float64
	balance float64
	left    *node
	right   *node
}

// Int type allows int comparison through Int method "Less"
type Int int

// Float type allows float comparison through Float method "Less"
type Float float64

// String type allows string comparison through String method "Less"
type String string

// Object type allows js.Value comparison through Object method "Less"
type Object js.Value

// Less returns true if a < b. Assume a, b are Int.
func (a Int) Less(b Interface) bool {
	return a < b.(Int)
}

// Less returns true if a < b. Assume a, b are Float.
func (a Float) Less(b Interface) bool {
	return a < b.(Float)
}

// Less returns true if a < b. Assume a, b are String.
func (a String) Less(b Interface) bool {
	return a < b.(String)
}

// Less returns true if a is less than b by whatever metric the user
// has defined for this datatype. Assume a, b are Object.
func (a Object) Less(b Interface) bool {
	return js.Value(a).Call("Less", (js.Value(b.(Object)))).Bool()
}

// checkType uses a type switch to determine the type of data being passed in.
// Given a built-in type, the data will be type casted to its nearest equivalent
// which implements Interface (e.g. int -> Int). Given a user-defined type, the
// data will be expected to have implemented a method "Less", and therefore
// implement Interface. If a type can implement Interface, it will be upcasted
// to the Interface type upon return. A user-defined type without a "Less"
// method will trigger a panic, exiting the program.
//
// Javascript types are handled identically, though they must first be converted
// into their Golang equivalents.
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
			}
			return Float(goVal.Float())
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

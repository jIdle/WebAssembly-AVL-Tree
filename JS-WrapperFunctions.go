package avl

import (
	"fmt"
	"reflect"
	"syscall/js"
)

/*	The following methods grant web programs access to this library.
	Once compiled into a WebAssembly module, these methods will be
	immediately available to call.
	(e.g. A Javascript program could query the AVL tree just like any
	other imported library.)

	To learn how to load and run a WebAssembly module within a Javascript
	program, have a look at the (amazing) documentation:
	https://developer.mozilla.org/en-US/docs/WebAssembly/Loading_and_running
*/

// InsertJS is a wrapper function for the AVL tree Insert.
func (t *AVL) InsertJS(this js.Value, args []js.Value) interface{} {
	t.Insert(args[0])
	return nil
}

// RemoveJS is a wrapper function for the AVL tree Remove.
func (t *AVL) RemoveJS(this js.Value, args []js.Value) interface{} {
	return t.Remove(args[0])
}

// RetrieveJS is a wrapper function for the AVL tree Retrieve.
func (t *AVL) RetrieveJS(this js.Value, args []js.Value) interface{} {
	switch v := t.Retrieve(args[0]).(type) {
	case Int:
		return int(v)
	case Float:
		return float64(v)
	case String:
		return string(v)
	case Object:
		return js.Value(v)
	default:
		panic(fmt.Sprintf("Could not find a compatible type conversion for %v\n", reflect.TypeOf(v).Name()))
	}
}

// AscendingJS is a wrapper function for the AVL tree Ascending traversal.
func (t *AVL) AscendingJS(this js.Value, args []js.Value) interface{} {
	return revertTypes(t.Ascending())
}

// DescendingJS is a wrapper function for the AVL tree Descending traversal.
func (t *AVL) DescendingJS(this js.Value, args []js.Value) interface{} {
	return revertTypes(t.Descending())
}

// PreorderJS is a wrapper function for the AVL tree Preorder traversal.
func (t *AVL) PreorderJS(this js.Value, args []js.Value) interface{} {
	return revertTypes(t.Preorder())
}

// PostorderJS is a wrapper function for the AVL tree Postorder traversal.
func (t *AVL) PostorderJS(this js.Value, args []js.Value) interface{} {
	return revertTypes(t.Postorder())
}

// LevelOrderJS is a wrapper function for the AVL tree LevelOrder traversal.
func (t *AVL) LevelOrderJS(this js.Value, args []js.Value) interface{} {
	array := t.LevelOrder()
	for i, e := range array {
		array[i] = revertTypes(e.([]interface{}))
	}
	return array
}

// revertTypes accepts an array of the user-defined types (necessary for generic
// storage withing the AVL tree) as an argument, and returns the same array but
// with each element having been converted to its equivalent built-in type.
// A non-user-defined type will trigger a panic, exiting the program.
func revertTypes(array []interface{}) []interface{} {
	for i, e := range array {
		switch v := e.(type) {
		case Int:
			array[i] = int(v)
		case Float:
			array[i] = float64(v)
		case String:
			array[i] = string(v)
		case Object:
			array[i] = js.Value(v)
		default:
			panic(fmt.Sprintf("Could not find a compatible type conversion for %v\n", reflect.TypeOf(v).Name()))
		}
	}
	return array
}

package main

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

// InsertJS : Wrapper function for AVL insert.
func (t *AVL) InsertJS(this js.Value, args []js.Value) interface{} {
	t.Insert(args[0])
	return nil
}

// RemoveJS : Wrapper function for AVL remove.
func (t *AVL) RemoveJS(this js.Value, args []js.Value) interface{} {
	return t.Remove(args[0])
}

// RetrieveJS : Wrapper function for AVL retrieve.
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

// AscendingJS : Wrapper function for AVL Ascending traversal.
func (t *AVL) AscendingJS(this js.Value, args []js.Value) interface{} {
	return revertTypes(t.Ascending())
}

// DescendingJS : Wrapper function for AVL Descending traversal.
func (t *AVL) DescendingJS(this js.Value, args []js.Value) interface{} {
	return revertTypes(t.Descending())
}

// PreorderJS : Wrapper function for AVL Preorder traversal.
func (t *AVL) PreorderJS(this js.Value, args []js.Value) interface{} {
	return revertTypes(t.Preorder())
}

// PostorderJS : Wrapper function for AVL Postorder traversal.
func (t *AVL) PostorderJS(this js.Value, args []js.Value) interface{} {
	return revertTypes(t.Postorder())
}

// LevelOrderJS : Wrapper function for AVL LevelOrder traversal.
func (t *AVL) LevelOrderJS(this js.Value, args []js.Value) interface{} {
	array := t.LevelOrder()
	for i, e := range array {
		array[i] = revertTypes(e.([]interface{}))
	}
	return array
}

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

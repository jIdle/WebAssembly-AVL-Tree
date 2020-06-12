package main

import "syscall/js"

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
}

// RemoveJS : Wrapper function for AVL remove.
func (t *AVL) RemoveJS(this js.Value, args []js.Value) interface{} {
	return t.Remove(args[0])
}

// RetrieveJS : Wrapper function for AVL retrieve.
func (t *AVL) RetrieveJS(this js.Value, args []js.Value) interface{} {
	return t.Retrieve(args[0])
}

// LevelOrderJS : Wrapper function for AVL LevelOrder.
func (t *AVL) LevelOrderJS(this js.Value, args []js.Value) interface{} {
	return t.LevelOrder(false)
}

// PreorderJS : Wrapper function for AVL Preorder.
func (t *AVL) PreorderJS(this js.Value, args []js.Value) interface{} {
	return t.Preorder()
}

// AscendingJS : Wrapper function for AVL Ascending.
func (t *AVL) AscendingJS(this js.Value, args []js.Value) interface{} {
	return t.Ascending()
}

// DescendingJS : Wrapper function for AVL Descending.
func (t *AVL) DescendingJS(this js.Value, args []js.Value) interface{} {
	return t.Descending()
}

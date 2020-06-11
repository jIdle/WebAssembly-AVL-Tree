package main

import (
	"fmt"
	"reflect"
)

// Interface : Allows comparison between keys in generic form.
type Interface interface {
	Less(toCompare Interface) bool
}

// Tree : Container for root node.
type Tree struct {
	root *node
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

// Remove custom types (e.g. Int, Float, String) after the test code has been fixed to not be absolute garbage.
func checkType(toConvert interface{}) Interface {
	switch t := toConvert.(type) {
	case int:
		return Int(t)
	case Int:
		return t
	case float64:
		return Float(t)
	case Float:
		return t
	case string:
		return String(t)
	case String:
		return t
	default:
		panic(fmt.Sprintf("Could not find a compatible type conversion for %v\n", reflect.TypeOf(t).Name()))
	}
}

package main

import "testing"

func Test_Tree_RemoveRotation(t *testing.T) {
	testTree := &Tree{root: nil}

	input := []int{1, 4, 4, 1}
	purpose := []string{
		"Single Left rotation",
		"Single Right rotation",
		"Left-Right rotation",
		"Right-Left rotation",
	}

	testTree.root = &node{data: Int(2), height: 3, balance: -1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(1), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: Int(3), height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right.right = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.right.right == nil {
		t.Logf("'%v': Deletion & rotation occur\t(PASSED)", purpose[0])
		if testTree.root.data == Int(3) && testTree.root.left.data == Int(2) && testTree.root.right.data == Int(4) {
			t.Logf("'%v': Correct deletion & rotation\t(PASSED)", purpose[0])
		} else {
			t.Errorf("'%v': Correct deletion & rotation\t(FAILED)", purpose[0])
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t(FAILED)", purpose[0])
	}

	testTree.root = &node{data: Int(3), height: 3, balance: 1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(2), height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.left.left = &node{data: Int(1), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[1])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.left.left == nil {
		t.Logf("'%v': Deletion & rotation occur \t(PASSED)", purpose[1])
		if testTree.root.data == Int(2) && testTree.root.left.data == Int(1) && testTree.root.right.data == Int(3) {
			t.Logf("'%v': Correct deletion & rotation\t(PASSED)", purpose[1])
		} else {
			t.Errorf("'%v': Correct deletion & rotation\t(FAILED)", purpose[1])
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur \t(FAILED)", purpose[1])
	}

	testTree.root = &node{data: Int(3), height: 3, balance: 1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(1), height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.left.right = &node{data: Int(2), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[2])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.left.right == nil {
		t.Logf("'%v': Deletion & rotation occur\t(PASSED)", purpose[2])
		if testTree.root.data == Int(2) && testTree.root.left.data == Int(1) && testTree.root.right.data == Int(3) {
			t.Logf("'%v': Correct deletion & rotation \t(PASSED)", purpose[2])
		} else {
			t.Errorf("'%v': Correct deletion & rotation \t(FAILED)", purpose[2])
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t(FAILED)", purpose[2])
	}

	testTree.root = &node{data: Int(2), height: 3, balance: -1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(1), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: Int(4), height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right.left = &node{data: Int(3), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[3])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.right.left == nil {
		t.Logf("'%v': Deletion & rotation occur\t(PASSED)", purpose[3])
		if testTree.root.data == Int(3) && testTree.root.left.data == Int(2) && testTree.root.right.data == Int(4) {
			t.Logf("'%v': Correct deletion & rotation \t(PASSED)", purpose[3])
		} else {
			t.Errorf("'%v': Correct deletion & rotation \t(FAILED)", purpose[3])
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t(FAILED)", purpose[3])
	}
}

func Test_Tree_InsertRotation(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
	}
	purpose := []string{
		"Single Left rotation",
		"Single Right rotation",
		"Left-Right rotation",
		"Right-Left rotation",
	}

	for i := 0; i < 4; i++ {
		testTree := &Tree{root: nil}
		for _, v := range inputs[i] {
			testTree.Insert(v)
		}
		if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil {
			t.Logf("'%v': Nodes created\t\t(PASSED)", purpose[i])
			if testTree.root.data == Int(2) && testTree.root.left.data == Int(1) && testTree.root.right.data == Int(3) {
				t.Logf("'%v': Correct rotation\t(PASSED)", purpose[i])
			} else {
				t.Errorf("'%v': Correct rotation\t(FAILED)", purpose[i])
			}
		} else {
			t.Errorf("'%v': Nodes created\t\t(FAILED)", purpose[i])
		}
	}
}

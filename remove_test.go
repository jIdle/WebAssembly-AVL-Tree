package main

import "testing"

func Test_Tree_Remove(t *testing.T) {
	testTree := &Tree{root: nil}

	input := []int{5, 5, 4, 5, 6, 5}
	purpose := []String{
		"Remove last node",
		"Remove w/ left subtree",
		"Remove at left subtree",
		"Remove w/ right subtree",
		"Remove at right subtree",
		"Remove & replace w/ IOS",
	}

	testTree.root = &node{data: Int(5), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[0])
	if testTree.root == nil {
		t.Logf("'%v': Node deleted\t\t(PASSED)", purpose[0])
	} else {
		t.Errorf("'%v': Node deleted\t\t(FAILED)", purpose[0])
	}

	testTree.root = &node{data: Int(5), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.left = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[1])
	if testTree.root != nil && testTree.root.left == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", purpose[1])
		if testTree.root.data == Int(4) {
			t.Logf("'%v': Node replaced\t(PASSED)", purpose[1])
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", purpose[1])
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", purpose[1])
	}

	testTree.root = &node{data: Int(5), height: 2, balance: 1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[2])
	if testTree.root != nil && testTree.root.left == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", purpose[2])
		if testTree.root.data == Int(5) {
			t.Logf("'%v': Node replaced\t(PASSED)", purpose[2])
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", purpose[2])
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", purpose[2])
	}

	testTree.root = &node{data: Int(5), height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right = &node{data: Int(6), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[3])
	if testTree.root != nil && testTree.root.right == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", purpose[3])
		if testTree.root.data == Int(6) {
			t.Logf("'%v': Node replaced\t(PASSED)", purpose[3])
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", purpose[3])
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", purpose[3])
	}

	testTree.root = &node{data: Int(5), height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right = &node{data: Int(6), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[4])
	if testTree.root != nil && testTree.root.right == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", purpose[4])
		if testTree.root.data == Int(5) {
			t.Logf("'%v': Node replaced\t(PASSED)", purpose[4])
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", purpose[4])
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", purpose[4])
	}

	testTree.root = &node{data: Int(5), height: 3, balance: -1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: Int(6), height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right.left = &node{data: Int(5), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(input[5])
	if testTree.root != nil && testTree.root.right != nil && testTree.root.left != nil && testTree.root.right.left == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", purpose[5])
		if testTree.root.data == Int(5) && testTree.root.left.data == Int(4) && testTree.root.right.data == Int(6) {
			t.Logf("'%v': Node replaced\t(PASSED)", purpose[5])
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", purpose[5])
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", purpose[5])
	}
}

package main

import "testing"

func Test_Tree_Insert(t *testing.T) {
	testTree := &Tree{root: nil}

	input := []int{5, 4, 6, 5}
	purpose := []string{
		"Insert into empty tree",
		"Insert less than",
		"Insert greater than",
		"Insert equal",
	}

	testTree.Insert(input[0])
	if testTree.root != nil {
		t.Logf("'%v': Node created\t(PASSED)", purpose[0])
		if testTree.root.data == Int(5) {
			t.Logf("'%v': Correct value\t(PASSED)", purpose[0])
		} else {
			t.Errorf("'%v': Correct value\t(FAILED)", purpose[0])
		}
	} else {
		t.Errorf("'%v': Node created\t(FAILED)", purpose[0])
	}

	testTree.Insert(input[1])
	if testTree.root != nil && testTree.root.left != nil {
		t.Logf("'%v': Node created\t\t(PASSED)", purpose[1])
		if testTree.root.data == Int(5) && testTree.root.left.data == Int(4) {
			t.Logf("'%v': Correct value\t(PASSED)", purpose[1])
		} else {
			t.Errorf("'%v': Correct value\t(FAILED)", purpose[1])
		}
	} else {
		t.Errorf("'%v': Node created\t\t(FAILED)", purpose[1])
	}

	testTree.Insert(input[2])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil {
		t.Logf("'%v': Node created\t(PASSED)", purpose[2])
		if testTree.root.data == Int(5) && testTree.root.left.data == Int(4) && testTree.root.right.data == Int(6) {
			t.Logf("'%v': Correct value\t(PASSED)", purpose[2])
		} else {
			t.Errorf("'%v': Correct value\t(FAILED)", purpose[2])
		}
	} else {
		t.Errorf("'%v': Node created\t(FAILED)", purpose[2])
	}

	testTree.Insert(input[3])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.right.left != nil {
		t.Logf("'%v': Node created\t\t(PASSED)", purpose[3])
		if testTree.root.data == Int(5) && testTree.root.left.data == Int(4) && testTree.root.right.data == Int(6) && testTree.root.right.left.data == Int(5) {
			t.Logf("'%v': Correct value\t\t(PASSED)", purpose[3])
		} else {
			t.Errorf("'%v': Correct value\t\t(FAILED)", purpose[3])
		}
	} else {
		t.Errorf("'%v': Node created\t\t(FAILED)", purpose[3])
	}
}

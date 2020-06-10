package main

import "testing"

type TestData struct {
	input   []Int
	purpose string
}

func Test_Tree_Search(t *testing.T) {
	testTree := &Tree{root: nil}
	tests := []TestData{
		{[]Int{4}, "Retrieve data at root"},
		{[]Int{2}, "Retrieve left child"},
		{[]Int{6}, "Retrieve right child"},
		{[]Int{3}, "Retrieve at left subtree"},
		{[]Int{5}, "Retrieve at right subtree"},
	}

	if _, e := testTree.Search(0); e != nil {
		t.Logf("'Retrieve from empty tree': Returned error\t(PASSED)")
	} else {
		t.Errorf("'Retrieve from empty tree': Returned error\t(FAILED)")
	}

	testTree.root = &node{data: Int(4), height: 3, balance: 0, left: nil, right: nil}
	testTree.root.left = &node{data: Int(2), height: 2, balance: -1, left: nil, right: nil}
	testTree.root.left.right = &node{data: Int(3), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: Int(6), height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right.left = &node{data: Int(5), height: 1, balance: 0, left: nil, right: nil}

	for i := 0; i < 5; i++ {
		if data, e := testTree.Search(tests[i].input[0]); e == nil {
			t.Logf("'%v': Value found  \t(PASSED)", tests[i].purpose)
			if data == tests[i].input[0] {
				t.Logf("'%v': Correct value\t(PASSED)", tests[i].purpose)
			} else {
				t.Errorf("'%v': Correct value\t(FAILED)", tests[i].purpose)
			}
		} else {
			t.Errorf("'%v': Value found  \t(FAILED)", tests[i].purpose)
		}
	}

	if _, e := testTree.Search(100); e != nil {
		t.Logf("'Value not stored': Returned error\t\t(PASSED)")
	} else {
		t.Errorf("'Value not stored': Returned error\t\t(FAILED)")
	}

}

func Test_Tree_Remove(t *testing.T) {
	testTree := &Tree{root: nil}
	tests := []TestData{
		{[]Int{5}, "Remove last node"},
		{[]Int{5}, "Remove w/ left subtree"},
		{[]Int{4}, "Remove at left subtree"},
		{[]Int{5}, "Remove w/ right subtree"},
		{[]Int{6}, "Remove at right subtree"},
		{[]Int{5}, "Remove & replace w/ IOS"},
	}

	testTree.root = &node{data: Int(5), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[0].input[0])
	if testTree.root == nil {
		t.Logf("'%v': Node deleted\t\t(PASSED)", tests[0].purpose)
	} else {
		t.Errorf("'%v': Node deleted\t\t(FAILED)", tests[0].purpose)
	}

	testTree.root = &node{data: Int(5), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.left = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[1].input[0])
	if testTree.root != nil && testTree.root.left == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[1].purpose)
		if testTree.root.data == Int(4) {
			t.Logf("'%v': Node replaced\t(PASSED)", tests[1].purpose)
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", tests[1].purpose)
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", tests[1].purpose)
	}

	testTree.root = &node{data: Int(5), height: 2, balance: 1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[2].input[0])
	if testTree.root != nil && testTree.root.left == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[2].purpose)
		if testTree.root.data == Int(5) {
			t.Logf("'%v': Node replaced\t(PASSED)", tests[2].purpose)
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", tests[2].purpose)
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", tests[2].purpose)
	}

	testTree.root = &node{data: Int(5), height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right = &node{data: Int(6), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[3].input[0])
	if testTree.root != nil && testTree.root.right == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[3].purpose)
		if testTree.root.data == Int(6) {
			t.Logf("'%v': Node replaced\t(PASSED)", tests[3].purpose)
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", tests[3].purpose)
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", tests[3].purpose)
	}

	testTree.root = &node{data: Int(5), height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right = &node{data: Int(6), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[4].input[0])
	if testTree.root != nil && testTree.root.right == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[4].purpose)
		if testTree.root.data == Int(5) {
			t.Logf("'%v': Node replaced\t(PASSED)", tests[4].purpose)
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", tests[4].purpose)
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", tests[4].purpose)
	}

	testTree.root = &node{data: Int(5), height: 3, balance: -1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: Int(6), height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right.left = &node{data: Int(5), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[5].input[0])
	if testTree.root != nil && testTree.root.right != nil && testTree.root.left != nil && testTree.root.right.left == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[5].purpose)
		if testTree.root.data == Int(5) && testTree.root.left.data == Int(4) && testTree.root.right.data == Int(6) {
			t.Logf("'%v': Node replaced\t(PASSED)", tests[5].purpose)
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", tests[5].purpose)
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", tests[5].purpose)
	}
}

func Test_Tree_RemoveRotation(t *testing.T) {
	testTree := &Tree{root: nil}
	tests := []TestData{
		{[]Int{1}, "Single Left rotation"},
		{[]Int{4}, "Single Right rotation"},
		{[]Int{4}, "Left-Right rotation"},
		{[]Int{1}, "Right-Left rotation"},
	}

	testTree.root = &node{data: Int(2), height: 3, balance: -1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(1), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: Int(3), height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right.right = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[0].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.right.right == nil {
		t.Logf("'%v': Deletion & rotation occur\t\t(PASSED)", tests[0].purpose)
		if testTree.root.data == Int(3) && testTree.root.left.data == Int(2) && testTree.root.right.data == Int(4) {
			t.Logf("'%v': Correct deletion & rotation\t(PASSED)", tests[0].purpose)
		} else {
			t.Errorf("'%v': Correct deletion & rotation\t(FAILED)", tests[0].purpose)
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t\t(FAILED)", tests[0].purpose)
	}

	testTree.root = &node{data: Int(3), height: 3, balance: 1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(2), height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.left.left = &node{data: Int(1), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[1].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.left.left == nil {
		t.Logf("'%v': Deletion & rotation occur \t(PASSED)", tests[1].purpose)
		if testTree.root.data == Int(2) && testTree.root.left.data == Int(1) && testTree.root.right.data == Int(3) {
			t.Logf("'%v': Correct deletion & rotation\t(PASSED)", tests[1].purpose)
		} else {
			t.Errorf("'%v': Correct deletion & rotation\t(FAILED)", tests[1].purpose)
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur \t(FAILED)", tests[1].purpose)
	}

	testTree.root = &node{data: Int(3), height: 3, balance: 1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(1), height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right = &node{data: Int(4), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.left.right = &node{data: Int(2), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[2].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.left.right == nil {
		t.Logf("'%v': Deletion & rotation occur\t\t(PASSED)", tests[2].purpose)
		if testTree.root.data == Int(2) && testTree.root.left.data == Int(1) && testTree.root.right.data == Int(3) {
			t.Logf("'%v': Correct deletion & rotation \t(PASSED)", tests[2].purpose)
		} else {
			t.Errorf("'%v': Correct deletion & rotation \t(FAILED)", tests[2].purpose)
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t\t(FAILED)", tests[2].purpose)
	}

	testTree.root = &node{data: Int(2), height: 3, balance: -1, left: nil, right: nil}
	testTree.root.left = &node{data: Int(1), height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: Int(4), height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right.left = &node{data: Int(3), height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[3].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.right.left == nil {
		t.Logf("'%v': Deletion & rotation occur\t\t(PASSED)", tests[3].purpose)
		if testTree.root.data == Int(3) && testTree.root.left.data == Int(2) && testTree.root.right.data == Int(4) {
			t.Logf("'%v': Correct deletion & rotation \t(PASSED)", tests[3].purpose)
		} else {
			t.Errorf("'%v': Correct deletion & rotation \t(FAILED)", tests[3].purpose)
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t\t(FAILED)", tests[3].purpose)
	}
}

func Test_Tree_Insert(t *testing.T) {
	testTree := &Tree{root: nil}
	tests := []TestData{
		{[]Int{5}, "Insert into empty tree"},
		{[]Int{4}, "Insert less than"},
		{[]Int{6}, "Insert greater than"},
		{[]Int{5}, "Insert equal"},
	}

	testTree.Insert(tests[0].input[0])
	if testTree.root != nil {
		t.Logf("'%v': Node created\t(PASSED)", tests[0].purpose)
		if testTree.root.data == Int(5) {
			t.Logf("'%v': Correct value\t(PASSED)", tests[0].purpose)
		} else {
			t.Errorf("'%v': Correct value\t(FAILED)", tests[0].purpose)
		}
	} else {
		t.Errorf("'%v': Node created\t(FAILED)", tests[0].purpose)
	}

	testTree.Insert(tests[1].input[0])
	if testTree.root != nil && testTree.root.left != nil {
		t.Logf("'%v': Node created\t\t(PASSED)", tests[1].purpose)
		if testTree.root.data == Int(5) && testTree.root.left.data == Int(4) {
			t.Logf("'%v': Correct value\t\t(PASSED)", tests[1].purpose)
		} else {
			t.Errorf("'%v': Correct value\t\t(FAILED)", tests[1].purpose)
		}
	} else {
		t.Errorf("'%v': Node created\t\t(FAILED)", tests[1].purpose)
	}

	testTree.Insert(tests[2].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil {
		t.Logf("'%v': Node created\t(PASSED)", tests[2].purpose)
		if testTree.root.data == Int(5) && testTree.root.left.data == Int(4) && testTree.root.right.data == Int(6) {
			t.Logf("'%v': Correct value\t(PASSED)", tests[2].purpose)
		} else {
			t.Errorf("'%v': Correct value\t(FAILED)", tests[2].purpose)
		}
	} else {
		t.Errorf("'%v': Node created\t(FAILED)", tests[2].purpose)
	}

	testTree.Insert(tests[3].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.right.left != nil {
		t.Logf("'%v': Node created\t\t(PASSED)", tests[3].purpose)
		if testTree.root.data == Int(5) && testTree.root.left.data == Int(4) && testTree.root.right.data == Int(6) && testTree.root.right.left.data == Int(5) {
			t.Logf("'%v': Correct value\t\t(PASSED)", tests[3].purpose)
		} else {
			t.Errorf("'%v': Correct value\t\t(FAILED)", tests[3].purpose)
		}
	} else {
		t.Errorf("'%v': Node created\t\t(FAILED)", tests[3].purpose)
	}
}

func Test_Tree_InsertRotation(t *testing.T) {
	tests := []TestData{
		{[]Int{1, 2, 3}, "Single Left rotation"},
		{[]Int{3, 2, 1}, "Single Right rotation"},
		{[]Int{3, 1, 2}, "Left-Right rotation"},
		{[]Int{1, 3, 2}, "Right-Left rotation"},
	}

	for i := 0; i < 4; i++ {
		testTree := &Tree{root: nil}
		for _, v := range tests[i].input {
			testTree.Insert(v)
		}
		if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil {
			t.Logf("'%v': Nodes created\t(PASSED)", tests[i].purpose)
			if testTree.root.data == Int(2) && testTree.root.left.data == Int(1) && testTree.root.right.data == Int(3) {
				t.Logf("'%v': Correct rotation\t(PASSED)", tests[i].purpose)
			} else {
				t.Errorf("'%v': Correct rotation\t(FAILED)", tests[i].purpose)
			}
		} else {
			t.Errorf("'%v': Nodes created\t(FAILED)", tests[i].purpose)
		}
	}
}

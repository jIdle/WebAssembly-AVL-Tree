package main

import "testing"

type TestData struct {
	input   []int
	purpose string
}

func Test_Tree_Remove(t *testing.T) {
	testTree := &Tree{root: nil}
	tests := []TestData{
		{[]int{5}, "Remove last node"},
		{[]int{5}, "Remove w/ left subtree"},
		{[]int{4}, "Remove from left subtree"},
		{[]int{5}, "Remove w/ right subtree"},
		{[]int{6}, "Remove from right subtree"},
		{[]int{5}, "Remove & replace w/ IOS"},
	}

	testTree.root = &node{data: 5, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[0].input[0])
	if testTree.root == nil {
		t.Logf("'%v': Node deleted\t\t(PASSED)", tests[0].purpose)
	} else {
		t.Errorf("'%v': Node deleted\t\t(FAILED)", tests[0].purpose)
	}

	testTree.root = &node{data: 5, height: 1, balance: 0, left: nil, right: nil}
	testTree.root.left = &node{data: 4, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[1].input[0])
	if testTree.root != nil && testTree.root.left == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[1].purpose)
		if testTree.root.data == 4 {
			t.Logf("'%v': Node replaced\t(PASSED)", tests[1].purpose)
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", tests[1].purpose)
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", tests[1].purpose)
	}

	testTree.root = &node{data: 5, height: 2, balance: 1, left: nil, right: nil}
	testTree.root.left = &node{data: 4, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[2].input[0])
	if testTree.root != nil && testTree.root.left == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[2].purpose)
		if testTree.root.data == 5 {
			t.Logf("'%v': Node replaced\t(PASSED)", tests[2].purpose)
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", tests[2].purpose)
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", tests[2].purpose)
	}

	testTree.root = &node{data: 5, height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right = &node{data: 6, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[3].input[0])
	if testTree.root != nil && testTree.root.right == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[3].purpose)
		if testTree.root.data == 6 {
			t.Logf("'%v': Node replaced\t(PASSED)", tests[3].purpose)
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", tests[3].purpose)
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", tests[3].purpose)
	}

	testTree.root = &node{data: 5, height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right = &node{data: 6, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[4].input[0])
	if testTree.root != nil && testTree.root.right == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[4].purpose)
		if testTree.root.data == 5 {
			t.Logf("'%v': Node replaced\t(PASSED)", tests[4].purpose)
		} else {
			t.Errorf("'%v': Node replaced\t(FAILED)", tests[4].purpose)
		}
	} else {
		t.Errorf("'%v': Node deleted\t(FAILED)", tests[4].purpose)
	}

	testTree.root = &node{data: 5, height: 3, balance: -1, left: nil, right: nil}
	testTree.root.left = &node{data: 4, height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: 6, height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right.left = &node{data: 5, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[5].input[0])
	if testTree.root != nil && testTree.root.right != nil && testTree.root.left != nil && testTree.root.right.left == nil {
		t.Logf("'%v': Node deleted\t(PASSED)", tests[5].purpose)
		if testTree.root.data == 5 && testTree.root.left.data == 4 && testTree.root.right.data == 6 {
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
		{[]int{1}, "Single Left rotation"},
		{[]int{4}, "Single Right rotation"},
		{[]int{4}, "Left-Right rotation"},
		{[]int{1}, "Right-Left rotation"},
	}

	testTree.root = &node{data: 2, height: 3, balance: -1, left: nil, right: nil}
	testTree.root.left = &node{data: 1, height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: 3, height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right.right = &node{data: 4, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[0].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.right.right == nil {
		t.Logf("'%v': Deletion & rotation occur\t\t(PASSED)", tests[0].purpose)
		if testTree.root.data == 3 && testTree.root.left.data == 2 && testTree.root.right.data == 4 {
			t.Logf("'%v': Correct deletion & rotation\t(PASSED)", tests[0].purpose)
		} else {
			t.Errorf("'%v': Correct deletion & rotation\t(FAILED)", tests[0].purpose)
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t\t(FAILED)", tests[0].purpose)
	}

	testTree.root = &node{data: 3, height: 3, balance: 1, left: nil, right: nil}
	testTree.root.left = &node{data: 2, height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right = &node{data: 4, height: 1, balance: 0, left: nil, right: nil}
	testTree.root.left.left = &node{data: 1, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[1].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.left.left == nil {
		t.Logf("'%v': Deletion & rotation occur\t(PASSED)", tests[1].purpose)
		if testTree.root.data == 2 && testTree.root.left.data == 1 && testTree.root.right.data == 3 {
			t.Logf("'%v': Correct deletion & rotation\t(PASSED)", tests[1].purpose)
		} else {
			t.Errorf("'%v': Correct deletion & rotation\t(FAILED)", tests[1].purpose)
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t(FAILED)", tests[1].purpose)
	}

	testTree.root = &node{data: 3, height: 3, balance: 1, left: nil, right: nil}
	testTree.root.left = &node{data: 1, height: 2, balance: -1, left: nil, right: nil}
	testTree.root.right = &node{data: 4, height: 1, balance: 0, left: nil, right: nil}
	testTree.root.left.right = &node{data: 2, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[2].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.left.right == nil {
		t.Logf("'%v': Deletion & rotation occur\t\t(PASSED)", tests[2].purpose)
		if testTree.root.data == 2 && testTree.root.left.data == 1 && testTree.root.right.data == 3 {
			t.Logf("'%v': Correct deletion & rotation\t(PASSED)", tests[2].purpose)
		} else {
			t.Errorf("'%v': Correct deletion & rotation\t(FAILED)", tests[2].purpose)
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t\t(FAILED)", tests[2].purpose)
	}

	testTree.root = &node{data: 2, height: 3, balance: -1, left: nil, right: nil}
	testTree.root.left = &node{data: 1, height: 1, balance: 0, left: nil, right: nil}
	testTree.root.right = &node{data: 4, height: 2, balance: 1, left: nil, right: nil}
	testTree.root.right.left = &node{data: 3, height: 1, balance: 0, left: nil, right: nil}
	testTree.Remove(tests[3].input[0])
	if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil && testTree.root.right.left == nil {
		t.Logf("'%v': Deletion & rotation occur\t\t(PASSED)", tests[3].purpose)
		if testTree.root.data == 3 && testTree.root.left.data == 2 && testTree.root.right.data == 4 {
			t.Logf("'%v': Correct deletion & rotation\t(PASSED)", tests[3].purpose)
		} else {
			t.Errorf("'%v': Correct deletion & rotation\t(FAILED)", tests[3].purpose)
		}
	} else {
		t.Errorf("'%v': Deletion & rotation occur\t\t(FAILED)", tests[3].purpose)
	}
}

func Test_Tree_Insert(t *testing.T) {
	testTree := &Tree{root: nil}
	tests := []TestData{
		{[]int{5}, "Insert into empty tree"},
		{[]int{4}, "Insert less than"},
		{[]int{6}, "Insert greater than"},
		{[]int{5}, "Insert equal"},
	}

	testTree.Insert(tests[0].input[0])
	if testTree.root != nil {
		t.Logf("'%v': Node created\t(PASSED)", tests[0].purpose)
		if testTree.root.data == 5 {
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
		if testTree.root.data == 5 && testTree.root.left.data == 4 {
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
		if testTree.root.data == 5 && testTree.root.left.data == 4 && testTree.root.right.data == 6 {
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
		if testTree.root.data == 5 && testTree.root.left.data == 4 && testTree.root.right.data == 6 && testTree.root.right.left.data == 5 {
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
		{[]int{1, 2, 3}, "Single Left rotation"},
		{[]int{3, 2, 1}, "Single Right rotation"},
		{[]int{3, 1, 2}, "Left-Right rotation"},
		{[]int{1, 3, 2}, "Right-Left rotation"},
	}

	for i := 0; i < 4; i++ {
		testTree := &Tree{root: nil}
		for _, v := range tests[i].input {
			testTree.Insert(v)
		}
		if testTree.root != nil && testTree.root.left != nil && testTree.root.right != nil {
			t.Logf("'%v': Nodes created\t(PASSED)", tests[i].purpose)
			if testTree.root.data == 2 && testTree.root.left.data == 1 && testTree.root.right.data == 3 {
				t.Logf("'%v': Correct rotation\t(PASSED)", tests[i].purpose)
			} else {
				t.Errorf("'%v': Correct rotation\t(FAILED)", tests[i].purpose)
			}
		} else {
			t.Errorf("'%v': Nodes created\t(FAILED)", tests[i].purpose)
		}
	}
}

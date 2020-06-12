package main_test

import "testing"

func Test_Tree_Retrieve(t *testing.T) {
	testTree := &Tree{root: nil}

	input := []int{4, 2, 6, 3, 5}
	purpose := []string{
		"Retrieve data at the root",
		"Retrieve the left child",
		"Retrieve the right child",
		"Retrieve at left subtree",
		"Retrieve at right subtree",
	}

	if _, e := testTree.Retrieve(0); e != nil {
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
		if data, e := testTree.Retrieve(input[i]); e == nil {
			t.Logf("'%v': Value found  \t(PASSED)", purpose[i])
			if data == Int(input[i]) {
				t.Logf("'%v': Correct value\t(PASSED)", purpose[i])
			} else {
				t.Errorf("'%v': Correct value\t(FAILED)", purpose[i])
			}
		} else {
			t.Errorf("'%v': Value found  \t(FAILED)", purpose[i])
		}
	}

	if _, e := testTree.Retrieve(100); e != nil {
		t.Logf("'Value not stored': Returned error\t\t(PASSED)")
	} else {
		t.Errorf("'Value not stored': Returned error\t\t(FAILED)")
	}

}

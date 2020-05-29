package main

import "testing"

type TestData struct {
	input   int
	purpose string
}

func Test_Tree_Insert(t *testing.T) {
	testTree := &Tree{root: nil}
	tests := []TestData{
		{5, "Insert into empty tree"},
		{4, "Insert less than"},
		{6, "Insert greater than"},
		{5, "Insert equal"},
	}

	testTree.Insert(tests[0].input)
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

	testTree.Insert(tests[1].input)
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

	testTree.Insert(tests[2].input)
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

	testTree.Insert(tests[3].input)
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

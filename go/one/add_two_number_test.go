package one

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := InitListNode(342)
	l2 := InitListNode(465)
	l := AddTwoNumbers(l1, l2)
	r := GetIntFromListNode(l)
	if r != 807 {
		t.Errorf("Add(342, 465) failed. Got %d, expected 807.", r)
	}

	l1 = InitListNode(342)
	l2 = InitListNode(965)
	l = AddTwoNumbers(l1, l2)
	r = GetIntFromListNode(l)
	if r != 1307 {
		t.Errorf("Add(342, 965) failed. Got %d, expected 1307.", r)
	}
}

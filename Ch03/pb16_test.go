package Ch03

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	size := 5
	head := genLinkList(size)

	rev := reverseList(head)

	expect := "4 3 2 1 0 "
	get := printList(rev)

	if get != expect {
		t.Errorf("Expect: %s BUT get: %s\n", expect, get)
	}
}

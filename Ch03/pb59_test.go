package Ch03

import "testing"

func TestRemoveDupFromSortedList(t *testing.T) {
	h := GenLinkLisFromValues([]int{1, 2, 3, 3, 4, 5, 6, 7, 7, 7, 7, 7})

	got := PrintList(removeDupFromSortedList(h))
	x := "1 2 3 4 5 6 7 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestRemoveDupFromSortedList_removeAllButHead(t *testing.T) {
	h := GenLinkLisFromValues([]int{1, 1, 1, 1, 1, 1, 1})

	got := PrintList(removeDupFromSortedList(h))
	x := "1 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestRemoveDupFromSortedList_noRemove(t *testing.T) {
	h := GenLinkLisFromValues([]int{1, 2, 3})

	got := PrintList(removeDupFromSortedList(h))
	x := "1 2 3 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestRemoveDupFromSortedList_noRemove_singleEl(t *testing.T) {
	h := GenLinkLisFromValues([]int{1})

	got := PrintList(removeDupFromSortedList(h))
	x := "1 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestRemoveDupFromSortedList_nilList(t *testing.T) {

	got := removeDupFromSortedList(nil)

	if got != nil {
		t.Errorf("Expect nil")
	}
}

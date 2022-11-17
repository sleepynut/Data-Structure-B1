package Ch03

import "testing"

func TestMerge2SortedListsXxx(t *testing.T) {
	n1 := []int{1, 2, 3, 4, 5, 10}
	n2 := []int{3, 4, 6, 7, 12, 15, 20}

	h1 := GenLinkLisFromValues(n1)
	h2 := GenLinkLisFromValues(n2)

	merged := merge2SortedLists(h1, h2)
	expected := "1 2 3 3 4 4 5 6 7 10 12 15 20 "
	got := PrintList(merged)

	if expected != got {
		t.Errorf("Expected: %s\nZGot: %s\n", expected, got)
	}
}

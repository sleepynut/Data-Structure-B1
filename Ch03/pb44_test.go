package Ch03

import "testing"

func TestSplitEvenOddList(t *testing.T) {
	size := 4
	h, _ := GenLinkList(size)
	h = splitEvenOddList(h)

	x := "0 2 1 3 "
	got := PrintList(h)

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestSplitEvenOddList_noOdd(t *testing.T) {

	h := GenLinkLisFromValues([]int{2, 4, 10, 6})
	h = splitEvenOddList(h)

	x := "2 4 10 6 "
	got := PrintList(h)

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestSplitEvenOddList_oneOddSingleList(t *testing.T) {

	h := GenLinkLisFromValues([]int{1})
	h = splitEvenOddList(h)

	x := "1 "
	got := PrintList(h)

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestSplitEvenOddList_allOdd(t *testing.T) {

	h := GenLinkLisFromValues([]int{3, 1, 5, 11})
	h = splitEvenOddList(h)

	x := "3 1 5 11 "
	got := PrintList(h)

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

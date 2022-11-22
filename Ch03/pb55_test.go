package Ch03

import "testing"

func TestArrangeNodes_even(t *testing.T) {
	h := GenLinkLisFromValues([]int{1, 2, 3, 4, 5, 6})

	got := PrintList(rearrangeNodes(h))
	x := "1 6 2 5 3 4 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}

}

func TestArrangeNodes_odd(t *testing.T) {
	h := GenLinkLisFromValues([]int{1, 2, 3, 4, 5, 6, 7})

	got := PrintList(rearrangeNodes(h))
	x := "1 7 2 6 3 5 4 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}

}

func TestArrangeNodes_two(t *testing.T) {
	h := GenLinkLisFromValues([]int{1, 2})

	got := PrintList(rearrangeNodes(h))
	x := "1 2 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}

}

func TestArrangeNodes_one(t *testing.T) {
	h := GenLinkLisFromValues([]int{1})

	got := PrintList(rearrangeNodes(h))
	x := "1 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}

}

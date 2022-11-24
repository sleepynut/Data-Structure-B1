package Ch03

import "testing"

func TestFindDuplicate(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{1, 2, 3, 6, 6, 10})
	h2 := GenLinkLisFromValues([]int{4, 5, 8, 8, 9, 10, 10, 12, 13})

	got := PrintList(findDuplicate(h1, h2))
	x := "10 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestFindDuplicate_noDup(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{1, 2, 3, 6, 6, 10})
	h2 := GenLinkLisFromValues([]int{4, 11, 15, 17, 20})

	got := findDuplicate(h1, h2)

	if got != nil {
		t.Errorf("Expect nil")
	}
}

func TestFindDuplicate_eqLength(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{1, 2, 3, 6, 6, 10})
	h2 := GenLinkLisFromValues([]int{4, 5, 6, 6, 10, 10})

	got := PrintList(findDuplicate(h1, h2))
	x := "6 10 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestFindDuplicate_singleNotEqual(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{1})
	h2 := GenLinkLisFromValues([]int{4})

	got := findDuplicate(h1, h2)

	if got != nil {
		t.Errorf("Expect nil")
	}
}

func TestFindDuplicate_singleEqual(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{4})
	h2 := GenLinkLisFromValues([]int{4})

	got := PrintList(findDuplicate(h1, h2))
	x := "4 "

	if x != got {
		t.Errorf("Expect:%s\nGot: %s\n", x, got)
	}
}

func TestFindDuplicate_nil(t *testing.T) {

	h2 := GenLinkLisFromValues([]int{4})

	got := findDuplicate(nil, h2)
	// x := "4 "

	if got != nil {
		t.Errorf("Expect nil")
	}
}

package Ch03

import "testing"

func TestAddDigitFromTwoLists(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{1, 2, 3, 4, 5})
	h2 := GenLinkLisFromValues([]int{1, 2, 3, 4, 1})

	h3 := addDigitFromTwoLists(h1, h2)
	got := PrintList(h3)
	x := "2 4 6 8 6 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestAddDigitFromTwoLists_carry(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{1, 2, 7, 4, 5})
	h2 := GenLinkLisFromValues([]int{1, 2, 3, 4, 1})

	h3 := addDigitFromTwoLists(h1, h2)
	got := PrintList(h3)
	x := "2 5 0 8 6 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestAddDigitFromTwoLists_carryExtend(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{9, 8, 7, 4, 5})
	h2 := GenLinkLisFromValues([]int{1, 2, 3, 4, 1})

	h3 := addDigitFromTwoLists(h1, h2)
	got := PrintList(h3)
	x := "1 1 1 0 8 6 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestAddDigitFromTwoLists_diff(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{1, 2, 3, 4, 5})
	h2 := GenLinkLisFromValues([]int{1, 2, 3})

	h3 := addDigitFromTwoLists(h1, h2)
	got := PrintList(h3)
	x := "1 2 4 6 8 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestAddDigitFromTwoLists_diffCarry(t *testing.T) {
	h1 := GenLinkLisFromValues([]int{1, 2, 7, 4, 5})
	h2 := GenLinkLisFromValues([]int{3, 2, 3})

	h3 := addDigitFromTwoLists(h1, h2)
	got := PrintList(h3)
	x := "1 3 0 6 8 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestAddDigitFromTwoLists_nil(t *testing.T) {

	h2 := GenLinkLisFromValues([]int{1, 2, 3, 4, 1})

	h3 := addDigitFromTwoLists(nil, h2)
	got := PrintList(h3)
	x := "1 2 3 4 1 "

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

func TestAddDigitFromTwoLists_bothNil(t *testing.T) {

	h3 := addDigitFromTwoLists(nil, nil)

	if h3 != nil {
		t.Errorf("Expect nil")
	}
}

package Ch03

import "testing"

func TestRearrangeFromX(t *testing.T) {
	h := GenLinkLisFromValues([]int{2, 5, 4, 3, 6, 3})
	x := 4

	got := PrintList(rearrangeFromX(h, x))
	ex := "2 3 3 5 4 6 "

	if got != ex {
		t.Errorf("Expect: %s\nGot: %s\n", ex, got)
	}
}

func TestRearrangeFromX_single(t *testing.T) {
	h := GenLinkLisFromValues([]int{2})
	x := 2

	got := PrintList(rearrangeFromX(h, x))
	ex := "2 "

	if got != ex {
		t.Errorf("Expect: %s\nGot: %s\n", ex, got)
	}
}

func TestRearrangeFromX_nil(t *testing.T) {
	x := 2

	got := rearrangeFromX(nil, x)

	if got != nil {
		t.Errorf("Expect nil")
	}
}

func TestRearrangeFromX_noChange(t *testing.T) {
	h := GenLinkLisFromValues([]int{0, 1, 2, 3, 4, 5})
	x := 10

	got := PrintList(rearrangeFromX(h, x))
	ex := "0 1 2 3 4 5 "

	if got != ex {
		t.Errorf("Expect: %s\nGot: %s\n", ex, got)
	}
}

func TestRearrangeFromX_noChange2(t *testing.T) {
	h := GenLinkLisFromValues([]int{0, 1, 2, 3, 4, 5})
	x := -1

	got := PrintList(rearrangeFromX(h, x))
	ex := "0 1 2 3 4 5 "

	if got != ex {
		t.Errorf("Expect: %s\nGot: %s\n", ex, got)
	}
}

func TestRearrangeFromX_headChange(t *testing.T) {
	h := GenLinkLisFromValues([]int{10, 1, 2, 3, 4, 5})
	x := 6

	got := PrintList(rearrangeFromX(h, x))
	ex := "1 2 3 4 5 10 "

	if got != ex {
		t.Errorf("Expect: %s\nGot: %s\n", ex, got)
	}
}

func TestRearrangeFromX_lastChange(t *testing.T) {
	h := GenLinkLisFromValues([]int{0, 1, 2, 3, 4, 5})
	x := 5

	got := PrintList(rearrangeFromX(h, x))
	ex := "0 1 2 3 4 5 "

	if got != ex {
		t.Errorf("Expect: %s\nGot: %s\n", ex, got)
	}
}

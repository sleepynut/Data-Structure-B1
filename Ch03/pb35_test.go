package Ch03

import (
	"testing"
)

func TestSplitListToTwo(t *testing.T) {
	size := 10
	ori := GenLinkList(size)

	h1, h2 := SplitListToTwo(ori)

	s1 := PrintList(h1)
	s2 := PrintList(h2)

	x1 := "0 1 2 3 4 "
	x2 := "5 6 7 8 9 "

	if s1 != x1 || s2 != x2 {
		t.Errorf("Expect h1 of: %s\nbut get %s\n\n", x1, s1)
		t.Errorf("Expect h2 of: %s\nbut get %s\n\n", x2, s2)
	}
}

func TestSplitListToTwoEdgeNil(t *testing.T) {
	size := 0
	ori := GenLinkList(size)

	h1, h2 := SplitListToTwo(ori)

	if h1 != nil || h2 != nil {
		t.Errorf("Expected nil pointer on each split lists")
	}
}

func TestSplitListToTwoEdgeOne(t *testing.T) {
	size := 1
	ori := GenLinkList(size)

	h1, h2 := SplitListToTwo(ori)

	x1 := "0 "
	s1 := PrintList(h1)

	if h2 != nil {
		t.Errorf("Expect second split list to be nil")
	}

	if s1 != x1 {
		t.Errorf("Expected first split list: %s\nbut got: %s\n", x1, s1)
	}

}

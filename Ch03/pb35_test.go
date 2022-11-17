package Ch03

import (
	"testing"
)

func TestSplitListToTwo(t *testing.T) {
	size := 10
	ori := genLinkList(size)

	h1, h2 := splitListToTwo(ori)

	s1 := printList(h1)
	s2 := printList(h2)

	x1 := "0 1 2 3 4 "
	x2 := "5 6 7 8 9 "

	if s1 != x1 || s2 != x2 {
		t.Errorf("Expect h1 of: %s\nbut get %s\n\n", x1, s1)
		t.Errorf("Expect h2 of: %s\nbut get %s\n\n", x2, s2)
	}
}

package Ch03

import "testing"

func TestFindSqrtNode(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)

	got := findSqrtNode(h).Val
	x := 2

	if x != got {
		t.Errorf("\nExpect: %d\nGot: %d", x, got)
	}

}

func TestFindSqrtNode_size1(t *testing.T) {
	size := 1
	h, _ := GenLinkList(size)

	got := findSqrtNode(h).Val
	x := 0

	if x != got {
		t.Errorf("\nExpect: %d\nGot: %d", x, got)
	}

}

func TestFindSqrtNode_exact(t *testing.T) {
	// sqrt(n) is an integer

	size := 9
	h, _ := GenLinkList(size)

	got := findSqrtNode(h).Val
	x := 2

	if x != got {
		t.Errorf("\nExpect: %d\nGot: %d", x, got)
	}

}

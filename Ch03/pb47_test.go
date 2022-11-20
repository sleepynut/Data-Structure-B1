package Ch03

import "testing"

func TestFindModularNodeFromLast(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)
	k := 3

	x := 7
	got := findModularNodeFromLast(h, k).Val

	if x != got {
		t.Errorf("Expected: %d\nGot: %d\n", x, got)
	}
}

func TestFindModularNodeFromLast_k1(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)
	k := 1

	x := 9
	got := findModularNodeFromLast(h, k).Val

	if x != got {
		t.Errorf("Expected: %d\nGot: %d\n", x, got)
	}
}

func TestFindModularNodeFromLast_exceed(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)
	k := 11

	// x := 7
	got := findModularNodeFromLast(h, k)

	if got != nil {
		t.Errorf("Expected nil")
	}
}

func TestFindModularNodeFromLas_head(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)
	k := 10

	x := 0
	got := findModularNodeFromLast(h, k).Val

	if x != got {
		t.Errorf("Expected: %d\nGot: %d\n", x, got)
	}
}

func TestFindModularNodeFromLast_invalidK(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)
	k := 0

	// x := 7
	got := findModularNodeFromLast(h, k)

	if got != nil {
		t.Errorf("Expected nil")
	}
}

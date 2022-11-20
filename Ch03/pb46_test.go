package Ch03

import "testing"

func TestFindModularNode(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)

	k := 3
	x := 8
	got := findModularNode(h, k).Val

	if x != got {
		t.Errorf("Expect: %d\nGot: %d\n", x, got)
	}
}

func TestFindModularNode_exceed(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)

	k := 15

	got := findModularNode(h, k)

	if got != nil {
		t.Errorf("Expect nil")
	}
}

func TestFindModularNode_last(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)

	k := 10

	x := 9
	got := findModularNode(h, k).Val

	if got != x {
		t.Errorf("Expect: %d\nGot: %d\n", x, got)
	}
}

func TestFindModularNode_k1(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)

	k := 1

	x := 9
	got := findModularNode(h, k).Val

	if got != x {
		t.Errorf("Expect: %d\nGot: %d\n", x, got)
	}
}

func TestFindModularNode_invalidK(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)

	k := 0

	// x := nil
	got := findModularNode(h, k)

	if got != nil {
		t.Errorf("Expect nil")
	}
}

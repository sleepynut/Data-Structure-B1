package Ch03

import "testing"

func TestRmvNodeFromList(t *testing.T) {
	size := 10
	h, _ := GenLinkList(size)

	tmp := h
	del := tmp.Next.Next.Next.Next

	d := rmvNodeFromList(h, del)
	got := PrintList(d)
	x := "0 1 2 3 5 6 7 8 9 "

	if x != got {
		t.Errorf("Expected: %s\nGot: %s \n", x, got)
	}
}

func TestRmvNodeFromList_rmvHead(t *testing.T) {
	size := 3
	h, _ := GenLinkList(size)

	d := rmvNodeFromList(h, h)
	got := PrintList(d)
	x := "1 2 "

	if x != got {
		t.Errorf("Expected: %s\nGot: %s \n", x, got)
	}
}

func TestRmvNodeFromList_rmvHeadSingleList(t *testing.T) {
	size := 1
	h, _ := GenLinkList(size)

	d := rmvNodeFromList(h, h)

	if d != nil {
		t.Errorf("Expected nil head")
	}
}

func TestRmvNodeFromList_rmvLast(t *testing.T) {
	size := 3
	h, _ := GenLinkList(size)
	tmp := h
	del := tmp.Next.Next

	d := rmvNodeFromList(h, del)

	x := "0 1 "
	got := PrintList(d)

	if x != got {
		t.Errorf("Expected: %s\nGot: %s\n", x, got)
	}
}

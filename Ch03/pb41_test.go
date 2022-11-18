package Ch03

import "testing"

func TestPrintRandList(t *testing.T) {
	size := 11
	ms := map[int]int{
		2: 4,
		3: 4,
		5: 10,
		6: 9,
	}

	rh, _ := GenRandLinkList(size, ms)
	got := PrintRandList(rh, make(map[*RandListNode]bool))
	x := "0 1 2 4 5 10 6 9 7 8 3 "

	if x != got {
		t.Errorf("Expected: %s\nGot: %s\n", x, got)
	}
}

func TestPrintRandList_BackwardRand(t *testing.T) {
	size := 11
	ms := map[int]int{
		2: 4,
		3: 4,
		9: 5,
		6: 10,
	}

	rh, _ := GenRandLinkList(size, ms)
	got := PrintRandList(rh, make(map[*RandListNode]bool))
	x := "0 1 2 4 5 6 10 7 8 9 3 "

	if x != got {
		t.Errorf("Expected: %s\nGot: %s\n", x, got)
	}
}

func TestCloneRandList(t *testing.T) {
	size := 11
	ms := map[int]int{
		2: 4,
		3: 4,
		9: 5,
		6: 10,
	}

	rh, _ := GenRandLinkList(size, ms)
	ch := CloneRandList(rh)

	x := PrintRandList(rh, make(map[*RandListNode]bool))
	got := PrintRandList(ch, make(map[*RandListNode]bool))

	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}
}

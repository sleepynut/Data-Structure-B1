package Ch03

import "testing"

func TestReverseKthFromList(t *testing.T) {
	head, _ := GenLinkList(10)

	k := 3
	rev := ReverseKthFromList(head, k)

	expected := "2 1 0 5 4 3 8 7 6 9 "
	got := PrintList(rev)

	if expected != got {
		t.Errorf("Expect: %s\nGet: %s\n", expected, got)
	}

}

func TestReverseKthFromList_nil(t *testing.T) {
	rev := ReverseKthFromList(nil, 10)

	if rev != nil {
		t.Errorf("Expect nil value")
	}

}

func TestReverseKthFromList_largeK(t *testing.T) {
	k := 10
	head, _ := GenLinkList(5)
	rev := ReverseKthFromList(head, k)
	x := "0 1 2 3 4 "
	got := PrintList(rev)

	if x != got {
		t.Errorf("Expected: %s\nGot: %s\n", x, got)
	}
}

func TestReverseKthFromList_keq1(t *testing.T) {
	k := 1
	head, _ := GenLinkList(5)
	rev := ReverseKthFromList(head, k)
	x := "0 1 2 3 4 "
	got := PrintList(rev)

	if x != got {
		t.Errorf("Expected: %s\nGot: %s\n", x, got)
	}
}

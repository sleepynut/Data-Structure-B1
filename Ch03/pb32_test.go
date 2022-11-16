package Ch03

import "testing"

func TestMergePairListXxx(t *testing.T) {
	cases := 3
	c1 := genLinkList(0)
	c2 := genLinkList(3)
	c3 := genLinkList(6)

	expected := []string{"", "1 0 2 ", "1 0 3 2 5 4 "}
	heads := []*ListNode{c1, c2, c3}

	for i := 0; i < cases; i++ {
		got := printList(swapPairList(heads[i]))
		if got != expected[i] {
			t.Errorf("Expected: %s\nGot: %s\n", expected[i], got)
		}
	}
}

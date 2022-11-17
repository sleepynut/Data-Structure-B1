package Ch03

import "testing"

func TestMergePairListXxx(t *testing.T) {
	cases := 3
	c2 := GenLinkList(3)
	c1 := GenLinkList(0)
	c3 := GenLinkList(6)

	expected := []string{"", "1 0 2 ", "1 0 3 2 5 4 "}
	heads := []*ListNode{c1, c2, c3}

	for i := 0; i < cases; i++ {
		got := PrintList(swapPairList(heads[i]))
		if got != expected[i] {
			t.Errorf("Expected: %s\nGot: %s\n", expected[i], got)
		}
	}
}

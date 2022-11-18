package Ch03

import (
	"strconv"
	"testing"
)

func TestGenCircularList(t *testing.T) {
	size := 5
	head, _ := GenCircularList(size)
	got := ""

	for i := 0; i <= size; i++ {
		got += strconv.Itoa(head.Val) + " "
		head = head.Next
	}

	x := "1 2 3 4 5 1 "
	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}

}

func TestGenCircularList_n1(t *testing.T) {
	size := 1
	head, _ := GenCircularList(size)
	got := ""

	for i := 0; i <= size; i++ {
		got += strconv.Itoa(head.Val) + " "
		head = head.Next
	}

	x := "1 1 "
	if x != got {
		t.Errorf("Expect: %s\nGot: %s\n", x, got)
	}

}

func TestJosephousCircle_n5k2(t *testing.T) {
	size, k := 5, 2

	got := josephousCircle(size, k)
	x := 3

	if got != x {
		t.Errorf("Expect: %d\nGot: %d\n", x, got)
	}
}

func TestJosephousCircle_n7k3(t *testing.T) {
	size, k := 7, 3

	got := josephousCircle(size, k)
	x := 4

	if got != x {
		t.Errorf("Expect: %d\nGot: %d\n", x, got)
	}
}

func TestJosephousCircle_n5k1(t *testing.T) {
	size, k := 5, 1

	got := josephousCircle(size, k)
	x := 5

	if got != x {
		t.Errorf("Expect: %d\nGot: %d\n", x, got)
	}
}

func TestJosephousCircle_n1k1(t *testing.T) {
	size, k := 1, 1

	got := josephousCircle(size, k)
	x := -1

	if got != x {
		t.Errorf("Expect: %d\nGot: %d\n", x, got)
	}
}

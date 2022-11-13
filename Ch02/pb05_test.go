package ch02

import "testing"

func TestFindLargestArea(t *testing.T) {
	m := [][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 1},
	}

	got := FindLargestRegion(m)
	want := 5

	if got != want {
		t.Errorf("Expect %d, but got %d from FindLargestArea\n", want, got)
	}
}

package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1([]int{0, 3, 6})
	if output != 436 {
		t.Error("Invalid output:", output)
	}

	output = pt1([]int{1, 0, 15, 2, 10, 13})
	if output != 2 {
		t.Error("Invalid:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 1 {
		t.Error("Invalid output:", output)
	}
}

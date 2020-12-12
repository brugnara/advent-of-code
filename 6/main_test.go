package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	if output != 6585 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 3276 {
		t.Error("Invalid output:", output)
	}
}

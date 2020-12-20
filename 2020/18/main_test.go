package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	if output != 5374004645253 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 88782789402798 {
		t.Error("Invalid output:", output)
	}
}

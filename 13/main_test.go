package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	if output != 171 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 539746751134958 {
		t.Error("Invalid output:", output)
	}
}

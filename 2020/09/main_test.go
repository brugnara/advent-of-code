package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	if output != 29221323 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput(), 29221323)
	if output != 4389369 {
		t.Error("Invalid output:", output)
	}
}

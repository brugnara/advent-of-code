package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	// output := pt1([]string{"^>v<"})
	if output != 377891 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 14110788 {
		t.Error("Invalid output:", output)
	}
}

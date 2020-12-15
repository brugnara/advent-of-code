package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	// output := pt1([]string{"^>v<"})
	if output != 2565 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 2639 {
		t.Error("Invalid output:", output)
	}
}

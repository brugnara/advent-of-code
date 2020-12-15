package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	if output != 2070 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 24179327893504 {
		t.Error("Invalid output:", output)
	}
}

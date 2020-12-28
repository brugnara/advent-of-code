package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	if output != 220 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 439 {
		t.Error("Invalid output:", output)
	}
}

func TestToInt(t *testing.T) {
	output := toInt([]bool{true, false, false, true})
	if output != 9 {
		t.Error("Invalid output:", output)
	}
}

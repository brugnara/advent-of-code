package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	if output != 842 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 617 {
		t.Error("Invalid output:", output)
	}
}

func TestComputeRow(t *testing.T) {
	output := computeRow("FBFBBFFRLR")
	if output != 44 {
		t.Error("invalid output:", output)
	}
}

func TestComputeCol(t *testing.T) {
	output := computeCol("FBFBBFFRLR")
	if output != 5 {
		t.Error("invalid output:", output)
	}
}

func TestGetID(t *testing.T) {
	output := getID(44, 5)
	if output != 357 {
		t.Error("invalid output:", output)
	}
}

package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1("yzbqklnj")
	// output := pt1([]string{"^>v<"})
	if output != 282749 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2("yzbqklnj")
	if output != 9962624 {
		t.Error("Invalid output:", output)
	}
}

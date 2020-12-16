package main

import "testing"

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	// output := pt1([]string{"^>v<"})
	if output != 258 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 53 {
		t.Error("Invalid output:", output)
	}
}

func TestIsNice(t *testing.T) {
	if !isNice("ugknbfddgicrmopn") {
		t.Error("Expected true")
		return
	}
	if !isNice("aaa") {
		t.Error("Expected true for aaa")
		return
	}
	for _, s := range []string{
		"jchzalrnumimnmhp",
		"haegwjzuvuyypxyu",
		"dvszwmarrgswjxmb",
	} {
		if isNice(s) {
			t.Error("Expected false for:", s)
		}
	}
}

func TestIsNice2(t *testing.T) {
	if !isNice2("qjhvhtzxzqqjkmpb") {
		t.Error("Should be nice")
	}
	if isNice2("uurcxstgmygtbstg") {
		t.Error("Should not be nice")
	}

	if isNice2("ieodomkazucvgmuy") {
		t.Error("should not be nice")
	}
}

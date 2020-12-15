package main

import (
	"testing"
)

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	if output != 6386593869035 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 4288986482164 {
		t.Error("Invalid output:", output)
	}
}

func TestCombo1(t *testing.T) {
	output := combo("000000000000000000000000000000X1101X", []int{30, 35}, 0)
	if output[0] != "000000000000000000000000000000011010" {
		t.Error("invalid output")
		return
	}
	if output[1] != "000000000000000000000000000000011011" {
		t.Error("invalid output")
		return
	}
	if output[2] != "000000000000000000000000000000111010" {
		t.Error("invalid output")
		return
	}
	if output[3] != "000000000000000000000000000000111011" {
		t.Error("invalid output")
		return
	}
}

func TestCombo2(t *testing.T) {
	output := combo("00000000000000000000000000000001X0XX", []int{32, 34, 35}, 0)
	/*
		for _, o := range output {
			fmt.Println(o)
		}
	*/
	if len(output) != 8 {
		t.Error("invalid output")
		return
	}
}

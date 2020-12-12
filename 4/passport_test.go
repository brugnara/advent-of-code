package main

import "testing"

func TestNewPassport(t *testing.T) {
	p := NewPassport()

	if len(p.validEyeColors) != 7 {
		t.Error("Invalid validEyeColors set")
	}
}

func TestAddValue(t *testing.T) {
	p := NewPassport()

	p.AddField("byr", "2020")
	if p.byr != 2020 {
		t.Error("Invalid byr!")
		return
	}

	p.AddField("iyr", "2021")
	if p.iyr != 2021 {
		t.Error("Invalid iyr")
		return
	}

	p.AddField("iyr", "foobar")
	if p.iyr != 0 {
		t.Error("Invalid iyr, should be 0 with invalid input")
		return
	}

	p.AddField("eyr", "99")
	if p.eyr != 99 {
		t.Error("Invalid eyr")
		return
	}

	// string values:
	values := []string{
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, val := range values {
		p.AddField(val, "value: "+val)
	}

	if p.hgt != "value: hgt" {
		t.Error("Invalid hgt")
		return
	}

	if p.hcl != "value: hcl" {
		t.Error("invalid hcl")
		return
	}

	if p.ecl != "value: ecl" {
		t.Error("invalid ecl")
		return
	}

	if p.pid != "value: pid" {
		t.Error("invalid pid")
		return
	}
}

func TestIsValid(t *testing.T) {
	p := NewPassport()

	p.AddField("byr", "2000")
	p.AddField("iyr", "2015")
	p.AddField("eyr", "2025")
	p.AddField("hgt", "160cm")
	p.AddField("hcl", "#123456")
	p.AddField("ecl", "blu")
	p.AddField("pid", "012345678")
	if !p.IsValid() {
		t.Error("should be valid!")
	}
}

func TestIsValidByr(t *testing.T) {
	p := NewPassport()

	p.AddField("iyr", "2015")
	p.AddField("eyr", "2025")
	p.AddField("hgt", "160cm")
	p.AddField("hcl", "#123456")
	p.AddField("ecl", "blu")
	p.AddField("pid", "012345678")

	p.AddField("byr", "1919")
	if p.IsValid() {
		t.Error("1919 should be invalid")
		return
	}

	p.AddField("byr", "2003")
	if p.IsValid() {
		t.Error("2003 should be invalid")
		return
	}

	p.AddField("byr", "1920")
	if !p.IsValid() {
		t.Error("1920 should be valid!")
		return
	}

	p.AddField("byr", "2002")
	if !p.IsValid() {
		t.Error("2002 should be valid!")
		return
	}
}

// and continue...

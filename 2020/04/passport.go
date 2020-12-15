package main

import (
	"regexp"
	"strconv"
)

// Passport stores a passport
type Passport struct {
	// valid eyeColor
	validEyeColors []string

	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
}

// NewPassport returns a new passport{}
func NewPassport() Passport {
	p := Passport{}
	p.validEyeColors = []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}
	return p
}

// AddField populate the passport with the given field: value
func (p *Passport) AddField(field, value string) {
	switch field {
	case "byr":
		p.byr, _ = strconv.Atoi(value)
	case "iyr":
		p.iyr, _ = strconv.Atoi(value)
	case "eyr":
		p.eyr, _ = strconv.Atoi(value)
	case "hgt":
		p.hgt = value
	case "hcl":
		p.hcl = value
	case "ecl":
		p.ecl = value
	case "pid":
		p.pid = value
	}
}

// IsValid checks passport
func (p Passport) IsValid() bool {
	/*
	  cid (Country ID) - ignored, missing or not.
	*/

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if p.byr < 1920 || p.byr > 2002 {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if p.iyr < 2010 || p.iyr > 2020 {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if p.eyr < 2020 || p.eyr > 2030 {
		return false
	}

	/*
	  hgt (Height) - a number followed by either cm or in:
	    If cm, the number must be at least 150 and at most 193.
	    If in, the number must be at least 59 and at most 76.
	*/
	re := regexp.MustCompile(`^(\d+)(\w+)$`)
	match := re.FindStringSubmatch(p.hgt)
	if len(match) != 3 {
		return false
	}
	nr, _ := strconv.Atoi(match[1])
	switch match[2] {
	case "cm":
		if nr < 150 || nr > 193 {
			return false
		}
	case "in":
		if nr < 59 || nr > 76 {
			return false
		}
	default:
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	re = regexp.MustCompile(`^\#[0-9a-f]{6}$`)
	match = re.FindStringSubmatch(p.hcl)
	if len(match) != 1 {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	foundEyeColor := false
	for _, color := range p.validEyeColors {
		if color == p.ecl {
			foundEyeColor = true
			break
		}
	}
	if !foundEyeColor {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	re = regexp.MustCompile(`^\d{9}$`)
	match = re.FindStringSubmatch(p.pid)
	if len(match) != 1 {
		return false
	}

	return true
}

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// https://adventofcode.com/2020/day/4

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

func pt1(lines []string) int {
	mandatory := map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
	}

	current := map[string]bool{}
	re := regexp.MustCompile(`(\w+)\:.+`)

	valids := 0

	for _, line := range lines {
		if line == "" {
			// fmt.Println("new passport")
			// fmt.Println(current)
			if len(current) == len(mandatory) {
				valids++
			}
			current = map[string]bool{}
			continue
		}
		for _, sub := range strings.Split(line, " ") {
			match := re.FindStringSubmatch(sub)
			if mandatory[match[1]] {
				current[match[1]] = true
			}
		}
	}
	return valids
}

func pt2(lines []string) int {
	current := NewPassport()
	re := regexp.MustCompile(`(\w+)\:(.+)`)

	valids := 0

	for _, line := range lines {
		if line == "" {
			// fmt.Println("new passport")
			// fmt.Println(current)
			if current.IsValid() {
				valids++
			}
			current = NewPassport()
			continue
		}
		for _, sub := range strings.Split(line, " ") {
			match := re.FindStringSubmatch(sub)
			current.AddField(match[1], match[2])
		}
	}
	return valids
}

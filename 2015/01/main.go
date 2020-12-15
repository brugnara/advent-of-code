package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// https://adventofcode.com/2015/day/1

func main() {
	lines := getInput()
	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")[0]
}

func pt1(line string) int {
	hash := map[rune]int{}
	for _, c := range line {
		hash[c]++
	}
	return hash['('] - hash[')']
}

func pt2(line string) int {
	pos := 0
	for i, c := range line {
		switch c {
		case '(':
			pos++
		case ')':
			pos--
		}
		if pos < 0 {
			return i + 1
		}
	}
	return -1
}

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/2

func main() {
	lines := getInput()
	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pt1(lines []string) (ret int) {
	re := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)
	for _, line := range lines {
		if line == "" {
			break
		}
		match := re.FindStringSubmatch(line)
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		c, _ := strconv.Atoi(match[3])
		l1 := a * b
		l2 := b * c
		l3 := a * c
		mn := min(min(l1, l2), l3)
		//
		ret += 2*l1 + 2*l2 + 2*l3 + mn
	}
	return
}

func pt2(lines []string) (ret int) {
	re := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)
	for _, line := range lines {
		if line == "" {
			break
		}
		match := re.FindStringSubmatch(line)
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		c, _ := strconv.Atoi(match[3])
		p1 := 2 * a
		p2 := 2 * b
		p3 := 2 * c
		mn := p1 + p2 + p3 - max(max(p1, p2), p3)
		//
		ret += mn + a*b*c
	}
	return
}

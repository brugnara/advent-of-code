package main

import (
	"fmt"
	"io/ioutil"
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
	group := map[rune]bool{}
	value := 0
	for _, line := range lines {
		if line == "" {
			// fmt.Println("new group, previous was:", group)
			value += len(group)
			group = map[rune]bool{}
			continue
		}
		//
		for _, c := range line {
			group[c] = true
		}

	}
	return value
}

func pt2(lines []string) int {
	group := map[rune]int{}
	groupLen := 0
	value := 0
	for _, line := range lines {
		if line == "" {
			// fmt.Println("new group, previous was:", group)
			for _, v := range group {
				if v == groupLen {
					value++
				}
			}
			group = map[rune]int{}
			groupLen = 0
			continue
		}
		//
		groupLen++
		for _, c := range line {
			group[c]++
		}

	}
	return value
}

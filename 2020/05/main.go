package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// https://adventofcode.com/2020/day/5

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
	maxID := 0

	for _, line := range lines {
		row := computeRow(line)
		col := computeCol(line)
		maxID = max(getID(row, col), maxID)
	}

	return maxID
}

func pt2(lines []string) int {
	seats := [128][8]bool{}
	for _, line := range lines {
		row := computeRow(line)
		col := computeCol(line)
		if row != -1 && col != -1 {
			seats[row][col] = true
		}
	}
	// search the missing seat! The valid empty one is surrounded by filled rows
	// so when a valid seat is spotted, the first empty will be our seat
	foundOne := false
	for i := 0; i < len(seats); i++ {
		for j, seat := range seats[i] {
			if !seat {
				if foundOne {
					// fmt.Println("empty seat found:", i, j)
					return getID(i, j)
				}
			} else {
				foundOne = true
			}
		}
	}
	fmt.Println(seats)
	return -1
}

func isValid(pattern string) bool {
	return len(pattern) == 10
}

func computeRow(pattern string) int {
	if !isValid(pattern) {
		return -1
	}
	start := 0
	end := 127
	for i := 0; i < 8; i++ {
		chr := pattern[i]
		if chr == 'F' {
			end = (start + end) / 2
		} else {
			start = (start + end) / 2
		}
	}
	return end
}

func computeCol(pattern string) int {
	if !isValid(pattern) {
		return -1
	}
	start := 0
	end := 7
	for i := 7; i < 10; i++ {
		chr := pattern[i]
		if chr == 'L' {
			end = (start + end) / 2
		} else {
			start = (start + end) / 2
		}
	}
	return end
}

func getID(row, col int) int {
	return row*8 + col
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

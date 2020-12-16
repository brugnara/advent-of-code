package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/5

func main() {
	lines := getInput()
	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

func pt1(lines []string) (ret int) {
	board := map[int]map[int]bool{}
	for i := 0; i < 1000; i++ {
		board[i] = map[int]bool{}
	}
	re := regexp.MustCompile(`(\w+) (\d+),(\d+) through (\d+),(\d+)`)
	for _, line := range lines {
		if line == "" {
			break
		}
		match := re.FindStringSubmatch(line)
		// fmt.Println(match[1:])
		x1, _ := strconv.Atoi(match[2])
		y1, _ := strconv.Atoi(match[3])
		x2, _ := strconv.Atoi(match[4])
		y2, _ := strconv.Atoi(match[5])
		//
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				switch match[1] {
				case "on":
					board[j][i] = true
				case "off":
					board[j][i] = false
				case "toggle":
					board[j][i] = !board[j][i]
				}
			}
		}
	}
	ret = 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if board[j][i] {
				ret++
			}
		}
	}
	return
}

func pt2(lines []string) (ret int) {
	board := map[int]map[int]int{}
	for i := 0; i < 1000; i++ {
		board[i] = map[int]int{}
	}
	re := regexp.MustCompile(`(\w+) (\d+),(\d+) through (\d+),(\d+)`)
	for _, line := range lines {
		if line == "" {
			break
		}
		match := re.FindStringSubmatch(line)
		// fmt.Println(match[1:])
		x1, _ := strconv.Atoi(match[2])
		y1, _ := strconv.Atoi(match[3])
		x2, _ := strconv.Atoi(match[4])
		y2, _ := strconv.Atoi(match[5])
		//
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				switch match[1] {
				case "on":
					board[j][i]++
				case "off":
					board[j][i]--
					if board[j][i] < 0 {
						board[j][i] = 0
					}
				case "toggle":
					board[j][i] += 2
				}
			}
		}
	}
	ret = 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			ret += board[j][i]
		}
	}
	return
}

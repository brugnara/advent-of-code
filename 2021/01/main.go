package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput() (ret []int) {
	raw, _ := ioutil.ReadFile("./input")
	lines := strings.Split(string(raw), "\n")

	for _, line := range lines {
		// save line as int into ret
		value, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		ret = append(ret, value)
	}

	return
}

func pt1(lines []int) (count int) {
	for i, line := range lines[1:] {
		if lines[i] < line {
			count++
		}
	}

	return
}

func pt2(lines []int) int {
	sums := []int{}
	for i, line := range lines[:len(lines)-2] {
		sums = append(sums, line+lines[i+1]+lines[i+2])
	}

	return pt1(sums)
}

func main() {
	lines := getInput()
	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

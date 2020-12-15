package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/9

const preambleLen = 25
const weakNr = 29221323

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines, weakNr))
}

func getInput() (ret []uint64) {
	raw, _ := ioutil.ReadFile("./input.txt")
	for _, nr := range strings.Split(string(raw), "\n") {
		n, _ := strconv.ParseUint(nr, 10, 64)
		if n != 0 {
			ret = append(ret, n)
		}
	}

	return
}

func pt1(lines []uint64) uint64 {
	for i := preambleLen; i < len(lines)-1; i++ {
		if !isValid(lines[i-preambleLen:i], lines[i]) {
			return lines[i]
		}
	}
	return 0
}

func isValid(slice []uint64, target uint64) bool {
	// fmt.Println("checking slice:", slice, "with target:", target)
	ln := len(slice)
	for i := 0; i < ln-1; i++ {
		for j := i + 1; j < ln; j++ {
			if slice[i]+slice[j] == target {
				return true
			}
		}
	}
	return false
}

func sum(slice []uint64) (ret uint64) {
	for _, n := range slice {
		ret += n
	}
	return
}

func min(slice []uint64) (ret uint64) {
	ret = slice[0]
	for _, n := range slice[1:] {
		if n < ret {
			ret = n
		}
	}
	return
}

func max(slice []uint64) (ret uint64) {
	ret = slice[0]
	for _, n := range slice[1:] {
		if n > ret {
			ret = n
		}
	}
	return
}

func pt2(lines []uint64, target uint64) (ret uint64) {
	ln := len(lines)

	for i := 0; i < ln-1; i++ {
		for j := i + 1; j < ln; j++ {
			if sum(lines[i:j]) == target {
				// fmt.Println(i, j-1, lines[i], lines[j-1], lines[i:j])
				return min(lines[i:j]) + max(lines[i:j])
				// return lines[i] + lines[j]
			}
		}
	}

	return
}

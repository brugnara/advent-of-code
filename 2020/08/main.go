package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/8

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

func pt1(lines []string) (accumulator int) {
	re := regexp.MustCompile(`(\w+) (\-\d+|\+\d+)`)
	done := map[int]bool{}
	row := 0

	for {
		line := lines[row]
		match := re.FindStringSubmatch(line)

		// fmt.Println("match:", match, "row:", row, "acc:", accumulator)

		if len(match) == 0 || done[row] {
			break
		}
		done[row] = true

		if match[1] == "nop" {
			row++
			continue
		}

		value, _ := strconv.Atoi(match[2])

		if match[1] == "jmp" {
			row += value
			continue
		}

		if match[1] == "acc" {
			row++
			accumulator += value
			continue
		}

	}

	return
}

func pt2(lines []string) (accumulator int) {
	re := regexp.MustCompile(`(\w+) (\-\d+|\+\d+)`)
	done := map[int]bool{}
	row := 0
	ln := len(lines)

	jmpToSkip := 1
	jmpSkipped := 0

	for {
		line := lines[row]
		match := re.FindStringSubmatch(line)

		// fmt.Println("match:", match, "row:", row, "acc:", accumulator)

		if len(match) == 0 {
			fmt.Println("Done correctly!")
			break
		}

		if done[row] {
			fmt.Println("Duplicated command")
			if jmpToSkip > ln {
				fmt.Println("Too many attempts!")
				break
			}
			// reset everything, we will try to skip a different JMP
			row = 0
			jmpSkipped = 0
			jmpToSkip++
			accumulator = 0
			done = map[int]bool{}
			continue
		}

		done[row] = true
		op := match[1]

		if op == "nop" {
			row++
			continue
		}

		value, _ := strconv.Atoi(match[2])

		if op == "jmp" {
			jmpSkipped++
			if jmpSkipped == jmpToSkip {
				fmt.Println("Skipping on row #", row)
				row++
				continue
			}
			// convert a jmp to nop
			row += value
			continue
		}

		if op == "acc" {
			row++
			accumulator += value
			continue
		}

	}

	return
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type slope struct {
	right, down int
}

// https://adventofcode.com/2020/day/3

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
	cols := len(lines[0])

	trees := 0

	for i, line := range lines {
		if len(line) < cols {
			continue
		}
		if line[(i*3)%cols] == '#' {
			trees++
		}
	}
	return trees
}

func pt2(lines []string) int {
	rows := len(lines)
	cols := len(lines[0])

	slopes := []slope{
		slope{1, 1},
		slope{3, 1},
		slope{5, 1},
		slope{7, 1},
		slope{1, 2},
	}

	result := 1
	for _, slp := range slopes {
		trees := 0

		for i := 0; i*slp.down < rows; i++ {
			line := lines[i*slp.down]
			if len(line) < cols {
				continue
			}
			// fmt.Println("checking:", i*slp.down, (i*slp.right)%cols)
			if line[(i*slp.right)%cols] == '#' {
				trees++
			}
		}

		// fmt.Println("slope:", slp, "has:", trees, "trees")

		result *= trees
	}
	return result
}

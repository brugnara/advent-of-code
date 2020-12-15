package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// https://adventofcode.com/2015/day/3

type point struct {
	x, y int
}

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
	current := point{0, 0}
	hash := map[point]int{
		current: 1,
	}
	for _, command := range lines[0] {
		fmt.Println(string(command))
		switch command {
		case '>':
			current.x++
		case '<':
			current.x--
		case 'v', 'V':
			current.y++
		case '^':
			current.y--
		}
		fmt.Println(current)
		hash[current]++
	}
	fmt.Println(hash)
	return len(hash)
}

func pt2(lines []string) (ret int) {
	posSanta := point{0, 0}
	posRobot := point{0, 0}
	hash := map[point]int{
		point{0, 0}: 2,
	}
	for i, command := range lines[0] {
		isSanta := i%2 == 1
		fmt.Println(string(command))
		switch command {
		case '>':
			if isSanta {
				posSanta.x++
			} else {
				posRobot.x++
			}
		case '<':
			if isSanta {
				posSanta.x--
			} else {
				posRobot.x--
			}
		case 'v', 'V':
			if isSanta {
				posSanta.y++
			} else {
				posRobot.y++
			}
		case '^':
			if isSanta {
				posSanta.y--
			} else {
				posRobot.y--
			}
		}

		if isSanta {
			hash[posSanta]++
		} else {
			hash[posRobot]++
		}
	}
	fmt.Println(hash)
	return len(hash)
}

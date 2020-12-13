package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// https://adventofcode.com/2020/day/11

type matrix map[int]map[int]byte

type change struct {
	x, y  int
	state byte
}

func main() {
	lines := getInput()
	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() (ret matrix) {
	raw, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(raw), "\n")
	ret = matrix{}
	for i, line := range lines {
		if line == "" {
			continue
		}
		ret[i] = map[int]byte{}
		for j, chr := range line {
			ret[i][j] = byte(chr)
		}
	}
	return
}

func checkCell(x, y int, g matrix) (bool, byte) {
	if g[y][x] == '.' {
		return false, '.'
	}

	adjacents := 0

	/*
		# # #
		X O #
		# # #
	*/
	if g[y][x-1] == '#' {
		adjacents++
	}

	/*
		X # #
		# O #
		# # #
	*/
	if g[y-1][x-1] == '#' {
		adjacents++
	}

	/*
		# X #
		# O #
		# # #
	*/
	if g[y-1][x] == '#' {
		adjacents++
	}

	/*
		# # X
		# O #
		# # #
	*/
	if g[y-1][x+1] == '#' {
		adjacents++
	}

	/*
		# # #
		# O X
		# # #
	*/
	if g[y][x+1] == '#' {
		adjacents++
	}

	/*
		# # #
		# O #
		# # X
	*/
	if g[y+1][x+1] == '#' {
		adjacents++
	}

	/*
		# # #
		# O #
		# X #
	*/
	if g[y+1][x] == '#' {
		adjacents++
	}

	/*
		# # #
		# O #
		X # #
	*/
	if g[y+1][x-1] == '#' {
		adjacents++
	}

	switch g[y][x] {
	case 'L':
		if adjacents == 0 {
			return true, '#'
		}
	case '#':
		if adjacents >= 4 {
			return true, 'L'
		}
	}

	return false, g[y][x]
}

func checkCell2(x, y, w, h int, g matrix) (bool, byte) {
	if g[y][x] == '.' {
		return false, '.'
	}

	adjacents := 0

	/*
		# # #
		X O #
		# # #
	*/
	for i := x - 1; i >= 0; i-- {
		if g[y][i] == 'L' {
			break
		}
		if g[y][i] == '#' {
			adjacents++
			break
		}
	}

	/*
		X # #
		# O #
		# # #
	*/
	i := 1
	for {
		if y-i < 0 || x-i < 0 {
			break
		}
		if g[y-i][x-i] == 'L' {
			break
		}
		if g[y-i][x-i] == '#' {
			adjacents++
			break
		}
		i++
	}

	/*
		# X #
		# O #
		# # #
	*/
	for i := y - 1; i >= 0; i-- {
		if g[i][x] == 'L' {
			break
		}
		if g[i][x] == '#' {
			adjacents++
			break
		}
	}

	/*
		# # X
		# O #
		# # #
	*/
	i = 1
	for {
		if y-i < 0 || x+i >= w {
			break
		}
		if g[y-i][x+i] == 'L' {
			break
		}
		if g[y-i][x+i] == '#' {
			adjacents++
			break
		}
		i++
	}

	/*
		# # #
		# O X
		# # #
	*/
	for i := x + 1; i < w; i++ {
		if g[y][i] == 'L' {
			break
		}
		if g[y][i] == '#' {
			adjacents++
			break
		}
	}

	/*
		# # #
		# O #
		# # X
	*/
	i = 1
	for {
		if y+i >= h || x+i >= w {
			break
		}
		if g[y+i][x+i] == 'L' {
			break
		}
		if g[y+i][x+i] == '#' {
			adjacents++
			break
		}
		i++
	}

	/*
		# # #
		# O #
		# X #
	*/
	for i := y + 1; i < h; i++ {
		if g[i][x] == 'L' {
			break
		}
		if g[i][x] == '#' {
			adjacents++
			break
		}
	}

	/*
		# # #
		# O #
		X # #
	*/
	i = 1
	for {
		if y+i >= h || x-i < 0 {
			break
		}
		if g[y+i][x-i] == 'L' {
			break
		}
		if g[y+i][x-i] == '#' {
			adjacents++
			break
		}
		i++
	}

	switch g[y][x] {
	case 'L':
		if adjacents == 0 {
			return true, '#'
		}
	case '#':
		if adjacents >= 5 {
			return true, 'L'
		}
	}

	return false, g[y][x]
}

func countOccupied(g matrix) (ret int) {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			if g[y][x] == '#' {
				ret++
			}
		}
	}
	return
}

func pt1(g matrix) (ret int) {
	// fmt.Println(g)
	height := len(g)
	width := len(g[0])
	moves := 0

	for {
		newStates := []change{}

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				willChange, newState := checkCell(x, y, g)
				// fmt.Println(willChange, newState)
				if willChange {
					newStates = append(newStates, change{x, y, newState})
				}
			}
		}
		// fmt.Println(newStates)
		// moves > 100 as a safe for the for {}
		if len(newStates) == 0 || moves > 100 {
			break
		}

		// apply changes
		// fmt.Println("applying", len(newStates), "new status")
		for _, state := range newStates {
			g[state.y][state.x] = state.state
		}
		moves++
	}

	// count occupied seats!
	return countOccupied(g)
}

func pt2(g matrix) (ret int) {
	height := len(g)
	width := len(g[0])
	moves := 0

	for {
		// p(g)
		newStates := []change{}

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				willChange, newState := checkCell2(x, y, width, height, g)
				// fmt.Println(willChange, newState)
				if willChange {
					newStates = append(newStates, change{x, y, newState})
				}
			}
		}
		// fmt.Println(newStates)
		// moves > 100 as a safe for the for {}
		if len(newStates) == 0 || moves > 100 {
			fmt.Println("moves:", moves)
			break
		}

		// apply changes
		// fmt.Println("applying", len(newStates), "new status")
		for _, state := range newStates {
			g[state.y][state.x] = state.state
		}
		moves++
	}

	// count occupied seats!
	return countOccupied(g)
}

func p(g matrix) {
	fmt.Println("##")
	for y := 0; y < len(g); y++ {
		line := g[y]
		l := ""
		for i := 0; i < len(line); i++ {
			l += string(line[i])
		}
		fmt.Println(l)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// https://adventofcode.com/2020/day/17

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

type myLevel map[int]map[int]bool
type myLevel2 map[int]map[int]map[int]bool

type myLevels map[int]myLevel
type myLevels2 map[int]myLevel2

func newLevel(h int) myLevel {
	level := myLevel{}
	for i := 0; i < h; i++ {
		level[i] = map[int]bool{}
	}
	return level
}
func newLevel2(h int) myLevel2 {
	level := myLevel2{}
	for i := 0; i < h; i++ {
		level[i] = map[int]map[int]bool{}
	}
	return level
}

func pt1(lines []string) (ret int) {
	levels := myLevels{}
	levels[0] = newLevel(len(lines))
	for row, line := range lines {
		if line == "" {
			break
		}
		for col, c := range line {
			levels[0][row][col] = c == '#'
		}
	}
	for iteration := 0; iteration < 6; iteration++ {
		//
		minLevel := 0
		maxLevel := 0
		minRows := 0
		maxRows := 0
		minCells := 0
		maxCells := 0
		for levelZ, level := range levels {
			if levelZ < minLevel {
				minLevel = levelZ
			}
			if levelZ > maxLevel {
				maxLevel = levelZ
			}
			for row, rowLevel := range level {
				// rows
				if row < minRows {
					minRows = row
				}
				if row > maxRows {
					maxRows = row
				}
				for col := range rowLevel {
					// cells
					if col < minCells {
						minCells = col
					}
					if col > maxCells {
						maxCells = col
					}
				}
			}
		}
		// fmt.Println(minLevel, maxLevel, minCells, maxCells, minRows, maxRows)
		todo := []int{}
		// for index, level := range levels {
		for index := minLevel - 1; index <= maxLevel+1; index++ {
			if _, ok := levels[index]; !ok {
				levels[index] = myLevel{}
			}
			level := levels[index]
			// fmt.Println(index, level)
			for row := minRows - 1; row <= maxRows+1; row++ {
				if _, ok := levels[index][row]; !ok {
					levels[index][row] = map[int]bool{}
				}
				for col := minCells - 1; col <= maxCells+1; col++ {
					adj := countAdjacents(levels, index, row, col)
					// fmt.Println(adj)
					if level[row][col] {
						if adj == 2 || adj == 3 {
							continue
						}
						// level[row][col] = false
						todo = append(todo, index, row, col, 0)
					} else {
						if adj == 3 {
							// level[row][col] = true
							todo = append(todo, index, row, col, 1)
						}
					}
				}
			}
		}
		// fmt.Println(todo)
		for i := 0; i < len(todo); i += 4 {
			levels[todo[i+0]][todo[i+1]][todo[i+2]] = todo[i+3] == 1
		}
		// fmt.Println(levels)
	}
	return count(levels)
}

func count(levels myLevels) (ret int) {
	for _, level := range levels {
		for _, row := range level {
			for _, col := range row {
				if col {
					ret++
				}
			}
		}
	}
	return
}

func count2(levels myLevels2) (ret int) {
	for _, level := range levels {
		for _, row := range level {
			for _, col := range row {
				for _, col2 := range col {
					if col2 {
						ret++
					}
				}
			}
		}
	}
	return
}

func countAdjacents(levels myLevels, level, row, col int) (ret int) {
	// remove self from count
	if levels[level][row][col] {
		ret = -1
	}
	for z := -1; z <= 1; z++ {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if levels[level+z][row+y][col+x] {
					ret++
				}
			}
		}
	}
	return
}

func countAdjacents2(levels myLevels2, level, row, col, col2 int) (ret int) {
	// remove self from count
	if levels[level][row][col][col2] {
		ret = -1
	}
	for z := -1; z <= 1; z++ {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for w := -1; w <= 1; w++ {
					if levels[level+z][row+y][col+x][col2+w] {
						ret++
					}
				}
			}
		}
	}
	return
}

func pt2(lines []string) (ret int) {
	levels := myLevels2{}
	levels[0] = newLevel2(len(lines))
	for row, line := range lines {
		if line == "" {
			break
		}
		for col, c := range line {
			levels[0][row][col] = map[int]bool{0: c == '#'}
		}
	}
	for iteration := 0; iteration < 6; iteration++ {
		//
		minLevel := 0
		maxLevel := 0
		minRows := 0
		maxRows := 0
		minCells := 0
		maxCells := 0
		minCells2 := 0
		maxCells2 := 0
		for levelZ, level := range levels {
			if levelZ < minLevel {
				minLevel = levelZ
			}
			if levelZ > maxLevel {
				maxLevel = levelZ
			}
			for row, rowLevel := range level {
				// rows
				if row < minRows {
					minRows = row
				}
				if row > maxRows {
					maxRows = row
				}
				for col, colLevel := range rowLevel {
					// cells
					if col < minCells {
						minCells = col
					}
					if col > maxCells {
						maxCells = col
					}
					for col2 := range colLevel {
						if col2 < minCells2 {
							minCells2 = col2
						}
						if col2 > maxCells2 {
							maxCells2 = col2
						}
					}
				}
			}
		}
		// fmt.Println(minLevel, maxLevel, minCells, maxCells, minRows, maxRows)
		todo := []int{}
		// for index, level := range levels {
		for index := minLevel - 1; index <= maxLevel+1; index++ {
			if _, ok := levels[index]; !ok {
				levels[index] = myLevel2{}
			}
			level := levels[index]
			// fmt.Println(index, level)
			for row := minRows - 1; row <= maxRows+1; row++ {
				if _, ok := levels[index][row]; !ok {
					levels[index][row] = map[int]map[int]bool{}
				}
				for col := minCells - 1; col <= maxCells+1; col++ {
					if _, ok := levels[index][row][col]; !ok {
						levels[index][row][col] = map[int]bool{}
					}
					for col2 := minCells2 - 1; col2 <= maxCells2+1; col2++ {
						adj := countAdjacents2(levels, index, row, col, col2)
						// fmt.Println(adj)
						if level[row][col][col2] {
							if adj == 2 || adj == 3 {
								continue
							}
							// level[row][col] = false
							todo = append(todo, index, row, col, col2, 0)
						} else {
							if adj == 3 {
								// level[row][col] = true
								todo = append(todo, index, row, col, col2, 1)
							}
						}
					}
				}
			}
		}
		// fmt.Println(todo)
		for i := 0; i < len(todo); i += 5 {
			levels[todo[i+0]][todo[i+1]][todo[i+2]][todo[i+3]] = todo[i+4] == 1
		}
		// fmt.Println(levels)
	}
	return count2(levels)
}

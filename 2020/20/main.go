package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/20

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

type tile struct {
	lines   [][]bool
	borders map[int][]int
}

func pt1(lines []string) (ret int) {
	hash := map[int]tile{}
	re := regexp.MustCompile(`Tile (\d+):`)
	current := 0
	tmpLines := [][]bool{}
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 0 {
			if current != 0 {
				hash[current] = tile{tmpLines, map[int][]int{}}
			}
			current, _ = strconv.Atoi(match[1])
			tmpLines = [][]bool{}
			continue
		}
		if line == "" {
			continue
		}
		tmpLines = append(tmpLines, lineToBool(line))
	}
	hash[current] = tile{tmpLines, map[int][]int{}}
	for k, v := range hash {
		hash[k].borders[0] = []int{
			toInt(v.lines[0]),
			toInt(reverse(v.lines[0])),
		}
		tmp := []bool{
			v.lines[0][0],
			v.lines[1][0],
			v.lines[2][0],
			v.lines[3][0],
			v.lines[4][0],
			v.lines[5][0],
			v.lines[6][0],
			v.lines[7][0],
			v.lines[8][0],
			v.lines[9][0],
		}
		hash[k].borders[1] = []int{
			toInt(tmp),
			toInt(reverse(tmp)),
		}
		//
		tmp = []bool{
			v.lines[0][9],
			v.lines[1][9],
			v.lines[2][9],
			v.lines[3][9],
			v.lines[4][9],
			v.lines[5][9],
			v.lines[6][9],
			v.lines[7][9],
			v.lines[8][9],
			v.lines[9][9],
		}
		hash[k].borders[2] = []int{
			toInt(tmp),
			toInt(reverse(tmp)),
		}
		//
		hash[k].borders[3] = []int{
			toInt(v.lines[9]),
			toInt(reverse(v.lines[9])),
		}
	}
	// find borders with two disjoints
	joints := map[int]int{}
	for k1, v1 := range hash {
		for _, vBorder := range v1.borders {
			for k2, v2 := range hash {
				for _, vBorder2 := range v2.borders {
					if vBorder2[0] == vBorder[0] || vBorder2[0] == vBorder[1] || vBorder2[1] == vBorder[0] || vBorder2[1] == vBorder[1] {
						if k1 != k2 {
							joints[k1]++
						}
					}
				}
			}
		}
	}
	fmt.Println(joints)
	ret = 1
	for k, v := range joints {
		if v == 2 {
			ret *= k
		}
	}
	return
}

func toInt(bArr []bool) (ret int) {
	for i, b := range bArr {
		if b {
			shift := (len(bArr) - 2 - i)
			if shift < 0 {
				ret++
			} else {
				ret += 2 << shift
			}
		}
	}
	return
}

func reverse(bArr []bool) (ret []bool) {
	ret = []bool{}
	for i := len(bArr) - 1; i >= 0; i-- {
		ret = append(ret, bArr[i])
	}
	return ret
}

func lineToBool(s string) (ret []bool) {
	ret = make([]bool, 10)
	for i, c := range s {
		ret[i] = c == '#'
	}
	return
}

func pt2(lines []string) (ret int) {

	return
}

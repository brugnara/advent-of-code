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
	id    int
	lines [][]bool
	// borders converted from bin to int, in 2 ways (LtoR, RtoL)
	borders map[int][]int
}

type myHugeMap map[int]map[int]tile

func (t *tile) reduce(x, y int) {
	t.lines = t.lines[0:1]
}

func genHash(lines []string) map[int]tile {
	hash := map[int]tile{}
	re := regexp.MustCompile(`Tile (\d+):`)
	current := 0
	tmpLines := [][]bool{}
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 0 {
			if current != 0 {
				hash[current] = tile{current, tmpLines, map[int][]int{}}
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
	// add the last one
	hash[current] = tile{current, tmpLines, map[int][]int{}}

	// populate the "borders"
	for k, v := range hash {
		hash[k].borders[1] = []int{
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
		hash[k].borders[0] = []int{
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
	return hash
}

func pt1(lines []string) (ret int) {
	hash := genHash(lines)
	// find borders with two matches
	joints := getJoints(hash)
	ret = 1
	for k, v := range joints {
		if v == 2 {
			ret *= k
		}
	}
	return
}

func getJoints(hash map[int]tile) map[int]int {
	joints := map[int]int{}
	for k1, v1 := range hash {
		for _, vBorder := range v1.borders {
			for k2, v2 := range hash {
				if k2 == k1 {
					continue
				}
				for _, vBorder2 := range v2.borders {
					if vBorder2[0] == vBorder[0] || vBorder2[0] == vBorder[1] ||
						vBorder2[1] == vBorder[0] || vBorder2[1] == vBorder[1] {
						joints[k1]++
					}
				}
			}
		}
	}
	return joints
}

func getJoints2(hash map[int]tile) map[int]map[int]bool {
	joints := map[int]map[int]bool{}
	for k1, v1 := range hash {
		for _, vBorder := range v1.borders {
			for k2, v2 := range hash {
				if k2 == k1 {
					continue
				}
				for _, vBorder2 := range v2.borders {
					if vBorder2[0] == vBorder[0] || vBorder2[0] == vBorder[1] ||
						vBorder2[1] == vBorder[0] || vBorder2[1] == vBorder[1] {
						if _, ok := joints[k1]; !ok {
							joints[k1] = map[int]bool{}
						}
						joints[k1][k2] = true
					}
				}
			}
		}
	}
	return joints
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

func rotate(t tile, count int) tile {
	if count <= 0 {
		return t
	}

	ln := len(t.lines)
	if ln != len(t.lines[0]) {
		fmt.Println(ln, "!=", len(t.lines[0]))
		panic("Y u do diss")
	}

	// rotating in-place, exchanging vertex
	for offset := 0; offset < ln/2; offset++ {
		for i := offset; i < ln-1-offset; i++ {
			t.lines[offset][i], t.lines[i][ln-1-offset] = t.lines[i][ln-1-offset], t.lines[offset][i]
			t.lines[i][ln-1-offset], t.lines[ln-1-offset][ln-1-i] = t.lines[ln-1-offset][ln-1-i], t.lines[i][ln-1-offset]
			t.lines[ln-1-offset][ln-1-i], t.lines[ln-1-i][offset] = t.lines[ln-1-i][offset], t.lines[ln-1-offset][ln-1-i]
		}
	}

	// exchanging borders too
	t.borders[0], t.borders[1], t.borders[2], t.borders[3] =
		t.borders[1], t.borders[2], t.borders[3], t.borders[0]

	// then flip
	t.borders[0][0], t.borders[0][1] = t.borders[0][1], t.borders[0][0]
	t.borders[2][0], t.borders[2][1] = t.borders[2][1], t.borders[2][0]
	return rotate(t, count-1)
}

func flip(t tile, direction byte) tile {
	ln1 := len(t.lines)
	ln2 := len(t.lines[0])
	//
	if direction == 'H' {
		for i := 0; i < ln1; i++ {
			for j := 0; j < ln2/2; j++ {
				t.lines[i][j], t.lines[i][ln2-1-j] = t.lines[i][ln2-1-j], t.lines[i][j]
			}
		}
		// exchanging borders too
		t.borders[1][0], t.borders[1][1] = t.borders[1][1], t.borders[1][0]
		t.borders[3][0], t.borders[3][1] = t.borders[3][1], t.borders[3][0]
		t.borders[0], t.borders[2] = t.borders[2], t.borders[0]
	} else {
		for i := 0; i < ln2; i++ {
			for j := 0; j < ln1/2; j++ {
				t.lines[j][i], t.lines[ln2-1-j][i] = t.lines[ln2-1-j][i], t.lines[j][i]
			}
		}
		// exchanging borders too
		t.borders[0][0], t.borders[0][1] = t.borders[0][1], t.borders[0][0]
		t.borders[2][0], t.borders[2][1] = t.borders[2][1], t.borders[2][0]
		t.borders[1], t.borders[3] = t.borders[3], t.borders[1]
	}

	return t
}

func pt2(lines []string) (ret int) {
	hash := genHash(lines)
	fmt.Println(len(hash))
	hugeMap := myHugeMap{}
	first := 0
	joints := getJoints2(hash)
	for k, v := range joints {
		if len(v) == 2 {
			first = k
			break
		}
	}
	fmt.Println(first, joints[first], hugeMap)
	fmt.Println(hash[first].borders)
	for v := range joints[first] {
		fmt.Println(hash[v].borders)
	}
	// rebuild the big image, starting from the first vertex
	// at first, we need to rotate the image until borders 0 and 1 are on
	// Right and Top (we are looking but these are always 0 and 1..)
	noMatches := []int{}
	for j := 0; j < 4; j++ {
		b := hash[first].borders[j][0]
		found := false
		for v := range joints[first] {
			// fmt.Println(hash[v].borders)
			for i := 0; i < 4; i++ {
				if b == hash[v].borders[i][0] || b == hash[v].borders[i][1] {
					found = true
					break
				}
			}
		}

		if !found {
			noMatches = append(noMatches, j)
		}
	}
	fmt.Println(noMatches)
	/*
			1
		0		2
			3
	*/
	// 1697 is a great choice to start from!
	// current = 1697
	current := first
	fmt.Println(current)
	count := 0
	y := 0
	x := 1
	hugeMap[y] = map[int]tile{
		0: hash[current],
	}
	for {
		// for security reasons, always keep a break somewhere..
		if count > 250 {
			fmt.Println("Breaking!")
			break
		}
		count++
		//
		// fmt.Println(hash[current].borders[2])
		tmpTile, id := getMatchingTile(hugeMap[y][x-1], 2, &hash)
		if id == -1 {
			tmpTile, id = getMatchingTile(hugeMap[y][0], 3, &hash)
			if id == -1 {
				break
			}
			x = 0
			y++
			fmt.Println("New line! Prev has a len of:", len(hugeMap[y-1]), "and has y:", y)
			hugeMap[y] = map[int]tile{}
		}
		hugeMap[y][x] = tmpTile
		x++
	}
	fmt.Println(len(hugeMap), len(hugeMap[0]), "count:", count)

	// this will print the map before cleanup!
	printHugeMap(&hugeMap)

	// zip the map, clearing boundaries
	zipped := zipMap(&hugeMap)
	fmt.Println(zipped, len(zipped), len(strings.Split(zipped, "\n")))

	// find re!
	reLine1 := regexp.MustCompile(`..................#.`)
	reLine2 := regexp.MustCompile(`#....##....##....###`)
	reLine3 := regexp.MustCompile(`.#..#..#..#..#..#...`)

	match1 := reLine1.FindAllStringIndex(zipped, -1)
	match2 := reLine2.FindAllStringIndex(zipped, -1)
	match3 := reLine3.FindAllStringIndex(zipped, -1)

	fmt.Println(len(match1), len(match2), len(match3))

	fmt.Println(match2[1:])
	fmt.Println(match3[1:])
	return
}

func zipMap(hugeMap *myHugeMap) string {
	ret := ""
	h := len(*hugeMap)
	w := len((*hugeMap)[0])
	hh := len((*hugeMap)[0][0].lines)
	ww := len((*hugeMap)[0][0].lines[0])
	//
	for y := 0; y < h; y++ {
		for j := 0; j < hh; j++ {
			line := ""
			for x := 0; x < w; x++ {
				for i := 0; i < ww; i++ {
					if x > 0 && i == 0 {
						continue
					}
					if y > 0 && j == 0 {
						continue
					}
					chr := "."
					if (*hugeMap)[y][x].lines[j][i] {
						chr = "#"
					}
					line += chr
				}
			}
			if len(line) > 0 {
				ret += line + "\n"
			}
		}
	}
	return ret
}

func printHugeMap(hugeMap *myHugeMap) {
	h := len(*hugeMap)
	w := len((*hugeMap)[0])
	hh := len((*hugeMap)[0][0].lines)
	ww := len((*hugeMap)[0][0].lines[0])

	for y := 0; y < h; y++ {
		for j := 0; j < hh; j++ {
			for x := 0; x < w; x++ {
				for i := 0; i < ww; i++ {
					chr := "."
					if (*hugeMap)[y][x].lines[j][i] {
						chr = "#"
					}
					fmt.Print(chr)
				}
			}
			fmt.Println()
		}
	}

}

// see main_test.go # TestOutput placeholder
func printHugeMapIDs(hugeMap *myHugeMap) {
	h := len(*hugeMap)
	w := len((*hugeMap)[0])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print((*hugeMap)[y][x].id, " ")
		}
		fmt.Println()
	}

}

func getMatchingTile(t tile, border int, hash *map[int]tile) (tile, int) {
	borders := t.borders[border]
	for k, v := range *hash {
		if k == t.id {
			continue
		}
		for i, b := range v.borders {
			switch border {
			case 2:
				if borders[0] == b[0] {
					switch i {
					// no need to rotate
					case 0:
						return v, k
					case 1, 2:
						return flip(rotate(v, i), 'V'), k
					case 3:
						return rotate(v, 3), k
					}
				}
				if borders[0] == b[1] {
					switch i {
					case 0:
						return flip(v, 'V'), k
					case 3:
						return flip(rotate(v, 3), 'V'), k
					default:
						return rotate(v, i), k
					}
				}
			case 3:
				if borders[0] == b[0] {
					switch i {
					case 0:
						return flip(rotate(v, 3), 'H'), k
					case 1:
						return v, k
					case 2:
						return rotate(v, 1), k
					case 3:
						return flip(rotate(v, 2), 'H'), k
					}
				}
				if borders[0] == b[1] {
					switch i {
					case 0:
						return rotate(v, 3), k
					case 1:
						return flip(v, 'H'), k
					case 2:
						return flip(rotate(v, 1), 'H'), k
					case 3:
						return rotate(v, 2), k
					}
				}
			}
		}
	}
	return tile{}, -1
}

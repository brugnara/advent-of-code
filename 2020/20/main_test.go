package main

import (
	"testing"
)

func TestPt1(t *testing.T) {
	output := pt1(getInput())
	if output != 23386616781851 {
		t.Error("Invalid output:", output)
	}
}

func TestPt2(t *testing.T) {
	output := pt2(getInput())
	if output != 23386616781851 {
		t.Error("Invalid output:", output)
	}
}

func TestToInt(t *testing.T) {
	output := toInt([]bool{true, false, false, true})
	if output != 9 {
		t.Error("Invalid output:", output)
	}
}

func TestRotate(t *testing.T) {
	tmp := tile{
		1,
		[][]bool{
			[]bool{true, false, false, false},
			[]bool{true, true, false, false},
			[]bool{true, true, true, false},
			[]bool{true, true, false, true},
		},
		map[int][]int{},
	}
	tmp.borders[0] = []int{8, 1}
	tmp.borders[1] = []int{12, 3}
	tmp.borders[2] = []int{14, 7}
	tmp.borders[3] = []int{13, 11}
	//

	rotated := rotate(tmp, 1)
	if rotated.borders[0][0] != 3 || rotated.borders[0][1] != 12 {
		t.Error("Wrong")
	}
	if rotated.borders[1][0] != 14 || rotated.borders[1][1] != 7 {
		t.Error("Wrong")
	}
	if rotated.borders[2][0] != 11 || rotated.borders[2][1] != 13 {
		t.Error("Wrong")
	}
	if rotated.borders[3][0] != 8 || rotated.borders[3][1] != 1 {
		t.Error("Wrong")
	}

	for k, v := range [][]bool{
		[]bool{false, false, false, true},
		[]bool{false, false, true, false},
		[]bool{false, true, true, true},
		[]bool{true, true, true, true},
	} {
		for i, vv := range v {
			if rotated.lines[k][i] != vv {
				t.Error("Whops")
			}
		}
	}

}

func TestFlip(t *testing.T) {
	tmp := tile{
		1,
		[][]bool{
			[]bool{true, false, false, false},
			[]bool{true, true, false, false},
			[]bool{true, true, true, false},
			[]bool{true, true, false, true},
		},
		map[int][]int{},
	}
	tmp.borders[0] = []int{8, 1}
	tmp.borders[1] = []int{12, 3}
	tmp.borders[2] = []int{14, 7}
	tmp.borders[3] = []int{13, 11}
	//

	flipped := flip(tmp, 'H')
	for k, v := range [][]bool{
		[]bool{false, false, false, true},
		[]bool{false, false, true, true},
		[]bool{false, true, true, true},
		[]bool{true, false, true, true},
	} {
		for i, vv := range v {
			if flipped.lines[k][i] != vv {
				t.Error("Whops")
			}
		}
	}

	if tmp.borders[0][0] != 14 || tmp.borders[0][1] != 7 {
		t.Error("Wrong")
	}

	if tmp.borders[2][0] != 8 || tmp.borders[2][1] != 1 {
		t.Error("Wrong")
	}

	if tmp.borders[1][0] != 3 || tmp.borders[1][1] != 12 {
		t.Error("Wrong")
	}

	if tmp.borders[3][0] != 11 || tmp.borders[3][1] != 13 {
		t.Error("Wrong")
	}
}

func TestFlipped2(t *testing.T) {
	tmp := tile{
		1,
		[][]bool{
			[]bool{true, false, false, false},
			[]bool{true, true, false, false},
			[]bool{true, true, true, false},
			[]bool{true, true, false, true},
		},
		map[int][]int{},
	}
	tmp.borders[0] = []int{8, 1}
	tmp.borders[1] = []int{12, 3}
	tmp.borders[2] = []int{14, 7}
	tmp.borders[3] = []int{13, 11}

	flipped := flip(tmp, 'V')
	for k, v := range [][]bool{
		[]bool{true, true, false, true},
		[]bool{true, true, true, false},
		[]bool{true, true, false, false},
		[]bool{true, false, false, false},
	} {
		for i, vv := range v {
			if flipped.lines[k][i] != vv {
				t.Error("Whops")
			}
		}
	}

	if tmp.borders[0][0] != 1 || tmp.borders[0][1] != 8 {
		t.Error("Wrong")
	}

	if tmp.borders[2][0] != 7 || tmp.borders[2][1] != 14 {
		t.Error("Wrong")
	}

	if tmp.borders[1][0] != 13 || tmp.borders[1][1] != 11 {
		t.Error("Wrong")
	}

	if tmp.borders[3][0] != 12 || tmp.borders[3][1] != 3 {
		t.Error("Wrong")
	}
}

func TestOutput(t *testing.T) {
	/*
		After sorting, this is the resulting matrix.
		Can be outputted in 4 orientations by chance.
		1399 1637 3373 3169 3011 2473 3253 2203 3847 1193 1949 2731
		3727 2089 1249 1747 2777 1063 2039 2131 1699 1009 3803 2927
		1811 2797 2437 1429 1621 1307 3371 3719 1861 1069 3631 3593
		1889 3709 3433 3701 3797 2099 1471 2621 1097 3083 3863 1103
		2837 1753 3313 3527 2851 1361 2887 2677 1907 3079 1013 1117
		1409 2503 1297 2647 1559 3697 1091 3323 3319 3359 1439 1181
		2707 2659 1667 2801 2861 2477 2069 1787 2029 2179 1381 3613
		1229 3203 2551 2273 2767 2243 2293 3533 1597 1367 1129 2879
		2543 3943 3121 2617 3761 2843 1151 2113 2539 1583 1019 1451
		3643 1579 2129 2971 3041 2423 1831 1327 3517 1171 2347 1913
		1031 2953 2939 1231 1283 3989 1453 2609 1879 2557 1619 2531
		3607 1289 3347 2143 1567 3413 3691 3911 2593 2657 3217 1697
	*/
}

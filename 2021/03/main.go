package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput() (ret []string) {
	raw, _ := ioutil.ReadFile("./input")
	lines := strings.Split(string(raw), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		ret = append(ret, line)
	}

	return
}

func getZeroAndOneCount(lines []string, i int) (zeroCount, oneCount int) {
	for _, line := range lines {
		if line[i] == '0' {
			zeroCount++
		} else {
			oneCount++
		}
	}
	return
}

func pt1(lines []string) (product int) {
	lineLen := len(lines[0])
	gamma := ""
	epsilon := ""

	for i := 0; i < lineLen; i++ {
		zeroCount, oneCount := getZeroAndOneCount(lines, i)

		for _, line := range lines {
			if line[i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		if oneCount > zeroCount {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	return multiplyBinaries(gamma, epsilon)
}

func multiplyBinaries(a, b string) int {
	aDec, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		panic(err)
	}
	bDec, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(aDec, bDec)

	return int(aDec * bDec)
}

func getLinesWithTheCharInPosition(lines []string, char byte, position int) (ret []string) {
	for _, line := range lines {
		if line[position] == char {
			ret = append(ret, line)
		}
	}
	return
}

func pt2(input []string) (product int) {
	lines := make([]string, len(input))
	copy(lines, input)
	lineLen := len(lines[0])

	oxigen := ""
	for i := 0; i < lineLen; i++ {
		zeros, ones := getZeroAndOneCount(lines, i)
		if zeros > ones {
			lines = getLinesWithTheCharInPosition(lines, '0', i)
		} else {
			lines = getLinesWithTheCharInPosition(lines, '1', i)
		}
		if len(lines) == 1 {
			oxigen = lines[0]
		}
	}

	lines = make([]string, len(input))
	copy(lines, input)

	co2 := ""
	for i := 0; i < lineLen; i++ {
		zeros, ones := getZeroAndOneCount(lines, i)
		if zeros < ones {
			lines = getLinesWithTheCharInPosition(lines, '0', i)
		} else {
			lines = getLinesWithTheCharInPosition(lines, '1', i)
		}
		if len(lines) == 1 {
			co2 = lines[0]
		}
	}

	return multiplyBinaries(co2, oxigen)
}

func main() {
	lines := getInput()
	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

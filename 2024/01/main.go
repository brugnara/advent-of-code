package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func getInput() (ret [][]int) {
	raw, _ := ioutil.ReadFile("./input")
	lines := strings.Split(string(raw), "\n")

	for _, line := range lines {

		tuple := strings.Split(line, "   ")

		if len(tuple) != 2 {
			continue
		}

		// save line as int into ret
		value1, err1 := strconv.Atoi(tuple[0])
		value2, err2 := strconv.Atoi(tuple[1])

		if err1 != nil || err2 != nil {
			continue
		}

		ret = append(ret, []int{value1, value2})
	}

	return
}

func pt1(tuples [][]int) (count int) {
	leftNumbers := []int{}
	rightNumbers := []int{}

	for _, tuple := range tuples {
		leftNumbers = append(leftNumbers, tuple[0])
		rightNumbers = append(rightNumbers, tuple[1])
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	for i, leftNumber := range leftNumbers {
		diff := leftNumber - rightNumbers[i]
		if diff < 0 {
			diff = diff * -1
		}

		count += diff
	}

	return
}

func pt2(lines [][]int) (count int) {
	leftNumbers := []int{}
	timesRightOccured := make(map[int]int)

	for _, tuple := range lines {
		leftNumbers = append(leftNumbers, tuple[0])
		timesRightOccured[tuple[1]]++
	}

	for _, leftNumber := range leftNumbers {
		count += timesRightOccured[leftNumber] * leftNumber
	}

	return
}

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

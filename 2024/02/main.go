package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput() (ret [][]int) {
	raw, _ := ioutil.ReadFile("./input")
	lines := strings.Split(string(raw), "\n")

	for _, line := range lines {
		numbers := strings.Split(line, " ")

		if len(numbers) == 0 {
			continue
		}

		convertedNumbers := []int{}

		for _, number := range numbers {
			value, err := strconv.Atoi(number)

			if err != nil {
				continue
			}

			convertedNumbers = append(convertedNumbers, value)
		}

		if len(convertedNumbers) > 0 {
			ret = append(ret, convertedNumbers)
		}
	}

	return
}

func isIncreasing(line []int) (increasing bool) {
	increasing = true

	for i := 1; i < len(line) && increasing; i++ {
		if line[i] < line[i-1] {
			increasing = false
		}
	}

	return
}

func checkLine(line []int) (lineValid bool) {
	lineValid = true

	if isIncreasing(line) {
		for i := 1; i < len(line) && lineValid; i++ {
			diff := line[i] - line[i-1]
			if diff > 3 || diff < 1 {
				lineValid = false
			}
		}
	} else {
		for i := len(line) - 1; i > 0 && lineValid; i-- {
			diff := line[i-1] - line[i]
			if diff > 3 || diff < 1 {
				lineValid = false
			}
		}
	}

	return
}

func pt1(lines [][]int) (count int) {
	for _, line := range lines {
		if checkLine(line) {
			count++
		}
	}

	return
}

func pt2(lines [][]int) (count int) {
	for _, line := range lines {
		if checkLine(line) {
			count++
			continue
		}

		lineAlreadyValid := false

		for i := 0; i < len(line) && !lineAlreadyValid; i++ {
			subLine := append([]int{}, line[:i]...)
			subLine = append(subLine, line[i+1:]...)

			if checkLine(subLine) {
				count++
				lineAlreadyValid = true
			}
		}
	}
	return
}

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

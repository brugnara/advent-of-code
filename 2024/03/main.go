package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func getInput() (ret []string) {
	raw, _ := ioutil.ReadFile("./input")
	return strings.Split(string(raw), "\n")
}

func computeMul(tuples []string) (count int) {
	numbers := strings.Split(tuples[1], ",")

	n1, err1 := strconv.Atoi(numbers[0])
	n2, err2 := strconv.Atoi(numbers[1])

	if err1 != nil || err2 != nil {
		return 0
	}

	return n1 * n2
}

func pt1(lines []string) (count int) {
	re := regexp.MustCompile(`mul\((\d+,\d+)\)`)

	for _, line := range lines {
		for _, tuples := range re.FindAllStringSubmatch(line, -1) {
			count += computeMul(tuples)
		}
	}

	return
}

func pt2(lines []string) (count int) {
	do := true
	reDo := regexp.MustCompile(`^do\(\)`)
	reDont := regexp.MustCompile(`^don't\(\)`)
	reMul := regexp.MustCompile(`^mul\((\d+,\d+)\)`)

	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			subLine := line[i:]
			if reDo.MatchString(subLine) {
				do = true
				continue
			}

			if reDont.MatchString(subLine) {
				do = false
				continue
			}

			match := reMul.FindStringSubmatch(subLine)

			if len(match) == 0 || !do {
				continue
			}

			count += computeMul(match)

		}
	}
	return
}

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

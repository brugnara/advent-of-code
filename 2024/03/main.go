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

func pt1(lines []string) (count int) {
	re := regexp.MustCompile(`mul\((\d+,\d+)\)`)

	for _, line := range lines {
		for _, tuples := range re.FindAllStringSubmatch(line, -1) {
			numbers := strings.Split(tuples[1], ",")

			n1, err1 := strconv.Atoi(numbers[0])
			n2, err2 := strconv.Atoi(numbers[1])

			if err1 != nil || err2 != nil {
				continue
			}

			count += n1 * n2
		}
	}

	return
}

func pt2(lines []string) (count int) {
	return
}

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

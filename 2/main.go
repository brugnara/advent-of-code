package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	raw, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(raw), "\n")

	pt1(lines)
	pt2(lines)
}

func pt2(lines []string) {
	valids := 0

	for _, line := range lines {
		re := regexp.MustCompile(`(\d+)\-(\d+) (\w): (\w+)`)
		match := re.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}

		char := match[3][0]
		low, _ := strconv.Atoi(match[1])
		high, _ := strconv.Atoi(match[2])

		if match[4][low-1] == char && match[4][high-1] != char {
			valids++
			continue
		}

		if match[4][low-1] != char && match[4][high-1] == char {
			valids++
		}
	}

	fmt.Println("Valids:", valids)
}

func pt1(lines []string) {
	valids := 0

	for _, line := range lines {
		re := regexp.MustCompile(`(\d+)\-(\d+) (\w): (\w+)`)
		match := re.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		count := 0
		for _, c := range match[4] {
			if c == rune(match[3][0]) {
				count++
			}
		}
		low, _ := strconv.Atoi(match[1])
		high, _ := strconv.Atoi(match[2])
		if count >= low && count <= high {
			valids++
		}
	}
	fmt.Println("total valids are:", valids)
}

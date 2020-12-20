package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
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

func do(ops string) string {
	re := regexp.MustCompile(`(\d+) (.{1}) (\d+)`)
	re2 := regexp.MustCompile(`\(|\)`)
	for {
		match := re.FindStringSubmatch(ops)
		if len(match) == 0 {
			break
		}
		var val int
		v1, _ := strconv.Atoi(match[1])
		v2, _ := strconv.Atoi(match[3])
		switch match[2][0] {
		case '+':
			val = v1 + v2
		case '*':
			val = v1 * v2
		}
		ops = strings.Replace(ops, strings.Join(match[1:], " "), strconv.Itoa(val), 1)
	}
	return re2.ReplaceAllString(ops, "")
}

func do2(ops string) string {
	res := []*regexp.Regexp{
		regexp.MustCompile(`(\d+) \+ (\d+)`),
		regexp.MustCompile(`(\d+) \* (\d+)`),
	}
	re2 := regexp.MustCompile(`\(|\)`)
	// +
	for index, operator := range []rune{'+', '*'} {
		for {
			match := res[index].FindStringSubmatch(ops)
			if len(match) == 0 {
				break
			}
			v1, _ := strconv.Atoi(match[1])
			v2, _ := strconv.Atoi(match[2])
			var val int
			switch operator {
			case '+':
				val = v1 + v2
			case '*':
				val = v1 * v2
			}
			ops = strings.Replace(
				ops,
				strings.Join(match[1:], string([]rune{' ', operator, ' '})),
				strconv.Itoa(val),
				1,
			)
			fmt.Println(ops)
		}
	}
	return re2.ReplaceAllString(ops, "")
}

func pt1(lines []string) (ret int) {
	re1 := regexp.MustCompile(`(\([\s\+\d\*]+\))`)
	for _, line := range lines {
		if line == "" {
			break
		}

		for {
			match := re1.FindAllStringSubmatch(line, -1)
			if len(match) == 0 {
				break
			}
			line = re1.ReplaceAllStringFunc(line, do)
		}
		nr, _ := strconv.Atoi(do(line))
		ret += nr
	}
	return
}

func pt2(lines []string) (ret int) {
	re1 := regexp.MustCompile(`(\([\s\+\d\*]+\))`)
	for _, line := range lines {
		if line == "" {
			break
		}

		for {
			match := re1.FindAllStringSubmatch(line, -1)
			if len(match) == 0 {
				break
			}
			line = re1.ReplaceAllStringFunc(line, do2)
		}
		nr, _ := strconv.Atoi(do2(line))
		ret += nr
	}
	return
}

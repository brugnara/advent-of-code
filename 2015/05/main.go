package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// https://adventofcode.com/2015/day/5

func main() {
	lines := getInput()
	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

func pt1(lines []string) (ret int) {
	for _, line := range lines {
		if isNice(line) {
			ret++
		}
	}
	return
}

func isNice(s string) bool {
	reWovel := regexp.MustCompile(`a|e|i|o|u`)
	reXX := regexp.MustCompile(`aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz`)
	reForbidden := regexp.MustCompile(`ab|cd|pq|xy`)

	matchWovel := reWovel.FindAllString(s, -1)
	if len(matchWovel) < 3 {
		return false
	}

	matchXX := reXX.FindString(s)
	if matchXX == "" {
		return false
	}

	matchForbidden := reForbidden.FindString(s)
	return matchForbidden == ""
}

func isNice2(s string) bool {
	hash := map[string]int{}
	// checks for aa..aa and also "aaa"
	oneValid := false
	for i := 0; i < len(s)-1; i++ {
		subset := string(s[i : i+2])
		if index, ok := hash[subset]; ok {
			if index == i-1 {
				return false
			}
			oneValid = true
		}
		hash[subset] = i
	}
	if !oneValid {
		return false
	}
	//
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func pt2(lines []string) (ret int) {
	for _, line := range lines {
		if isNice2(line) {
			ret++
		}
	}
	return
}

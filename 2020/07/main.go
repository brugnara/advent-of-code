package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/7

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

func pt1(lines []string) (valids int) {
	re1 := regexp.MustCompile(`(.+) contain (.+)`)
	reColor := regexp.MustCompile(`^(.+) bags$`)
	reCnt := regexp.MustCompile(`(\d+) (.+) bag`)
	bags := map[string][]string{}

	for _, line := range lines {
		match := re1.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		// fmt.Println(match[1], " -> ", match[2])
		matchColor := reColor.FindStringSubmatch(match[1])
		// fmt.Println(matchColor[1])

		// find contains
		for _, color := range strings.Split(match[2], ", ") {
			matchCnt := reCnt.FindStringSubmatch(color)
			// fmt.Println(matchCnt)
			if len(matchCnt) > 0 {
				bags[matchColor[1]] = append(bags[matchColor[1]], matchCnt[2])
			}
		}
		//
	}
	// fmt.Println(bags)
	for bag := range bags {
		if hasGold(bag, bags) {
			valids++
		}
	}
	return
}

func pt2(lines []string) (valids int) {
	re1 := regexp.MustCompile(`(.+) contain (.+)`)
	reColor := regexp.MustCompile(`^(.+) bags$`)
	reCnt := regexp.MustCompile(`(\d+) (.+) bag`)

	bags := map[string][]bag{}

	for _, line := range lines {
		match := re1.FindStringSubmatch(line)
		if len(match) == 0 {
			continue
		}
		// fmt.Println(match[1], " -> ", match[2])
		matchColor := reColor.FindStringSubmatch(match[1])
		// fmt.Println(matchColor[1])

		// find contains
		for _, color := range strings.Split(match[2], ", ") {
			matchCnt := reCnt.FindStringSubmatch(color)
			// fmt.Println(matchCnt)
			if len(matchCnt) > 0 {
				count, _ := strconv.Atoi(matchCnt[1])
				bags[matchColor[1]] = append(bags[matchColor[1]], bag{matchCnt[2], count})
			}
		}
		//
	}
	return countInside("shiny gold", bags)
}

func hasGold(bag string, bags map[string][]string) bool {
	for _, b := range bags[bag] {
		if b == "shiny gold" {
			return true
		}
		if hasGold(b, bags) {
			return true
		}
	}
	return false
}

func countInside(bag string, bags map[string][]bag) (ret int) {
	if _, ok := bags[bag]; !ok {
		// fmt.Println(bag, "has nothing inside")
		return 0
	}
	// fmt.Println("Inside of", bag, "there are:", bags[bag])
	for _, inside := range bags[bag] {
		// fmt.Println("Inside #", i, "is:", inside.name, "with size:", inside.count)
		ret += inside.count + inside.count*countInside(inside.name, bags)
	}
	return
}

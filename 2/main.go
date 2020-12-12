package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	pt1()
}

func pt1() {
	raw, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(raw), "\n")

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
			fmt.Println("Valid:", line)
			valids++
		} else {
			fmt.Println("invalid:", line)
		}
	}
	fmt.Println("total valids are:", valids)
}

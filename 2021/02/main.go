package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type command struct {
	direction string
	amount    int
}

func getInput() (ret []command) {
	raw, _ := ioutil.ReadFile("./input")
	lines := strings.Split(string(raw), "\n")
	re := regexp.MustCompile(`(\w+) (\d+)`)

	for _, line := range lines {
		if line == "" {
			continue
		}

		matchMem := re.FindStringSubmatch(line)
		amount, _ := strconv.Atoi(matchMem[2])
		ret = append(ret, command{
			direction: matchMem[1],
			amount:    amount,
		})
	}

	return
}

func pt1(commands []command) (product int) {
	x := 0
	depth := 0

	for _, command := range commands {
		switch command.direction {
		case "down":
			depth += command.amount
		case "up":
			depth -= command.amount
		case "forward":
			x += command.amount
		}
	}

	return x * depth
}

func pt2(commands []command) (product int) {
	aim := 0
	depth := 0
	x := 0

	for _, command := range commands {
		switch command.direction {
		case "down":
			aim += command.amount
		case "up":
			aim -= command.amount
		case "forward":
			x += command.amount
			depth += aim * command.amount
		}
	}

	return x * depth
}

func main() {
	commands := getInput()
	fmt.Println("PT1:", pt1(commands))
	fmt.Println("PT2:", pt2(commands))
}

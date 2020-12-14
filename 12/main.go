package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/12

type position struct {
	x, y int
}

func main() {
	lines := getInput()
	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

func turn(where, currentDirection byte, degrees int) byte {
	/*
			   N
			 W + E
		  	 S
	*/
	compass := []byte{'N', 'E', 'S', 'W'}
	index := -1
	for i, c := range compass {
		if c == currentDirection {
			index = i
		}
	}
	if where == 'L' && degrees != 180 {
		degrees += 180
	}
	return compass[(index+degrees/90)%4]
}

func pt1(commands []string) (ret int) {
	x := 0
	y := 0
	dir := byte('E')
	/*
			N
		W	+ E
			S
	*/
	re := regexp.MustCompile(`(\w)(\d+)`)
	for _, cmd := range commands {
		match := re.FindStringSubmatch(cmd)
		if len(match) == 0 {
			continue
		}
		amount, _ := strconv.Atoi(match[2])
		action := match[1][0]
		switch action {
		case 'L', 'R':
			dir = turn(action, dir, amount)
		case 'F':
			switch dir {
			case 'N':
				x -= amount
			case 'S':
				x += amount
			case 'W':
				y -= amount
			case 'E':
				y += amount
			}
		case 'N':
			x -= amount
		case 'S':
			x += amount
		case 'W':
			y -= amount
		case 'E':
			y += amount
		}
	}
	// fmt.Println(x, y)
	return abs(x) + abs(y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func turnWaypoint(action byte, wp, ship position, amount int) position {
	if amount == 0 {
		return wp
	}
	if amount == 180 {
		return position{
			-wp.x,
			-wp.y,
		}
	}
	if (action == 'R' && amount == 90) || (action == 'L' && amount == 270) {
		return position{
			wp.y,
			-wp.x,
		}
	}
	if (action == 'L' && amount == 90) || (action == 'R' && amount == 270) {
		return position{
			-wp.y,
			wp.x,
		}
	}
	return wp
}

func pt2(commands []string) (ret int) {
	waypoint := position{-1, 10}
	ship := position{0, 0}
	/*
			N
		W	+ E
			S
	*/
	re := regexp.MustCompile(`(\w)(\d+)`)
	for _, cmd := range commands {
		// fmt.Println(waypoint, ship)
		match := re.FindStringSubmatch(cmd)
		if len(match) == 0 {
			break
		}
		amount, _ := strconv.Atoi(match[2])
		action := match[1][0]
		switch action {
		case 'L', 'R':
			// dir = turn(action, dir, amount)
			waypoint = turnWaypoint(action, waypoint, ship, amount)
		case 'F':
			ship.x += amount * waypoint.x
			ship.y += amount * waypoint.y
		case 'N':
			waypoint.x -= amount
		case 'S':
			waypoint.x += amount
		case 'W':
			waypoint.y -= amount
		case 'E':
			waypoint.y += amount
		}
	}
	// fmt.Println(ship.x, ship.y)
	return abs(ship.x) + abs(ship.y)
}

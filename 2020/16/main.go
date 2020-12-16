package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type slope struct {
	right, down int
}

// https://adventofcode.com/2020/day/16

type point struct {
	from, to int
}

type myData struct {
	info   map[string][]point
	ticket []int
	nearby [][]int
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

func decode(lines []string) myData {
	ret := myData{
		map[string][]point{},
		[]int{},
		[][]int{},
	}
	reHeader := regexp.MustCompile(`(.+): (\d+)-(\d+) or (\d+)-(\d+)`)

	for i := 0; i < 20; i++ {
		// fmt.Println(lines[i])
		match := reHeader.FindStringSubmatch(lines[i])
		fmt.Println(match[1:])
		tmp := make([]int, 4)
		for i := 2; i <= 5; i++ {
			tmp[i-2], _ = strconv.Atoi(match[i])
		}
		ret.info[match[1]] = []point{
			point{tmp[0], tmp[1]},
			point{tmp[2], tmp[3]},
		}
	}

	// myticket
	ret.ticket = convertToTicket(lines[22])

	// tickets
	i := 25
	for {
		line := lines[i]
		if line == "" {
			break
		}
		ret.nearby = append(ret.nearby, convertToTicket(line))
		i++
	}

	return ret
}

func convertToTicket(s string) []int {
	ret := make([]int, 20)
	for i, v := range strings.Split(s, ",") {
		ret[i], _ = strconv.Atoi(v)
	}
	return ret
}

func pt1(lines []string) (ret int) {
	data := decode(lines)
	// fmt.Println(data)
	for _, ticket := range data.nearby {
		for _, nr := range ticket {
			valid := false
			for _, v := range data.info {
				if (nr >= v[0].from && nr <= v[0].to) || (nr >= v[1].from && nr <= v[1].to) {
					valid = true
					break
				}
			}
			if !valid {
				ret += nr
			}
		}
	}
	return
}

func isValid(ticket []int, data myData) (valid bool) {
	for _, nr := range ticket {
		valid = false
		for _, v := range data.info {
			if (nr >= v[0].from && nr <= v[0].to) || (nr >= v[1].from && nr <= v[1].to) {
				valid = true
				break
			}
		}
		if !valid {
			return
		}
	}
	return
}

func guessHeader(data myData) map[string]int {
	validTickets := [][]int{}
	for _, ticket := range data.nearby {
		if isValid(ticket, data) {
			fmt.Println("VALID")
			validTickets = append(validTickets, ticket)
		} else {
			// fmt.Println("INVALID")
		}
	}
	fmt.Println(len(validTickets))
	guess := map[string]map[int]bool{}
	for k := range data.info {
		guess[k] = map[int]bool{}
	}
	for index := 0; index < 20; index++ {
		for k, v := range data.info {
			// fmt.Println("Guessing:", k, v)
			allValid := true
			for _, ticket := range validTickets {
				nr := ticket[index]
				if (nr >= v[0].from && nr <= v[0].to) || (nr >= v[1].from && nr <= v[1].to) {
					continue
				}
				allValid = false
				break
			}
			if allValid {
				guess[k][index] = true
			}
		}
	}
	fixGuess(&guess)
	ret := map[string]int{}
	for k, v := range guess {
		for kk := range v {
			ret[k] = kk
		}
	}
	return ret
}

func fixGuess(guess *map[string]map[int]bool) {
	edited := false
	for k, v := range *guess {
		if len(v) == 1 {
			// found a slot sure, clean every other
			key := 0
			for kk := range v {
				key = kk
			}
			for kk, vv := range *guess {
				if k != kk {
					if len(vv) > 1 {
						edited = true
					}
					delete(vv, key)
				}
			}
			break
		}
	}
	if !edited {
		return
	}
	fixGuess(guess)
}

func pt2(lines []string) int {
	data := decode(lines)
	guess := guessHeader(data)
	ret := 1
	for k, v := range guess {
		fmt.Println(k, v)
		if strings.HasPrefix(k, "departure") {
			ret *= data.ticket[v]
		}
	}
	return ret
}

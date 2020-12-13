package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/2

func main() {
	input := getInput()
	fmt.Println("PT1:", pt1(input))
	fmt.Println("PT2:", pt2(input))
}

func getInput() (ret []int) {
	raw, _ := ioutil.ReadFile("./input.txt")
	array := strings.Split(string(raw), "\n")
	for _, a := range array {
		if a == "" {
			continue
		}
		nr, _ := strconv.Atoi(a)
		ret = append(ret, nr)
	}
	sort.Ints(ret)
	return
}

func pt1(input []int) int {
	// count from 0 to 1st and one "3" for the final device
	hash := map[int]int{input[0]: 1, 3: 1}
	for i := 1; i < len(input); i++ {
		delta := input[i] - input[i-1]
		// fmt.Println(input[i], input[i-1], delta)
		hash[delta]++
	}
	// fmt.Println(hash)
	return hash[1] * hash[3]
}

func pt2(input []int) uint64 {
	hash := map[int]uint64{}
	hash[0] = 1
	max := 1
	for _, nr := range input {
		hash[nr] = hash[nr-1] + hash[nr-3] + hash[nr-2]
		if nr > max {
			max = nr
		}
	}
	return hash[max]
}

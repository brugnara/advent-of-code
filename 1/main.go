package main

import (
	"fmt"
	"io/ioutil"
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
	return
}

func pt1(input []int) int {
	for i := 0; i < len(input)-1; i++ {
		for j := i; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				// fmt.Println(input[i], input[j], input[i]+input[j], input[i]*input[j])
				return input[i] * input[j]
			}
		}
	}
	return -1
}

func pt2(input []int) int {
	// this is not so "smart" but the input is small enough to handle a "dumb"
	// solution, evolving in around O(n^3) ops
	for i := 0; i < len(input)-2; i++ {
		for j := i; j < len(input)-1; j++ {
			for z := j; z < len(input); z++ {
				if input[i]+input[j]+input[z] == 2020 {
					// fmt.Println(input[i], input[j], input[z], input[i]+input[j]+input[z], input[i]*input[j]*input[z])
					return input[i] * input[j] * input[z]
				}
			}
		}
	}
	return -1
}

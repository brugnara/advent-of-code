package main

import (
	"fmt"
)

// https://adventofcode.com/2020/day/2

func main() {
	lines := getInput()
	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []int {
	// return []int{1, 0, 15, 2, 10, 13}
	return []int{1, 0, 15, 2, 10, 13}
}

func pt1(nums []int) int {
	hash := map[int][]int{}
	history := []int{}
	current := 0
	for i, n := range nums {
		hash[n] = []int{i}
		history = append(history, n)
		current++
	}
	fmt.Println(hash, history)
	for current < 2020 {
		prev := history[current-1]
		// fmt.Println("Prev:", prev)
		var next = 0
		ln := len(hash[prev])
		if ln != 1 {
			next = hash[prev][ln-1] - hash[prev][ln-2]
		}
		if _, ok := hash[next]; !ok {
			hash[next] = []int{}
		}
		// fmt.Println("Next", next)
		hash[next] = append(hash[next], current)
		history = append(history, next)
		current++
	}
	// fmt.Println(current, len(history))
	// fmt.Println(history)
	return history[current-1]
}

func pt2(nums []int) int {
	hash := map[int][]int{}
	history := []int{}
	current := 0
	for i, n := range nums {
		hash[n] = []int{i}
		history = append(history, n)
		current++
	}
	fmt.Println(hash, history)
	for current < 30000000 {
		prev := history[current-1]
		// fmt.Println("Prev:", prev)
		var next = 0
		ln := len(hash[prev])
		if ln != 1 {
			next = hash[prev][ln-1] - hash[prev][ln-2]
		}
		if _, ok := hash[next]; !ok {
			hash[next] = []int{}
		}
		// fmt.Println("Next", next)
		hash[next] = append(hash[next], current)
		history = append(history, next)
		current++
	}
	// fmt.Println(current, len(history))
	// fmt.Println(history)
	return history[current-1]
}

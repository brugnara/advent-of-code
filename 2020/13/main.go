package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/13

type park struct {
	ID     uint64
	offset uint64
}

func main() {
	fmt.Println("PT1:", pt1(getInput()))
	fmt.Println("PT2:", pt2(getInput()))
}

func getInput() (ts int, busses []int) {
	raw, _ := ioutil.ReadFile("./input.txt")
	array := strings.Split(string(raw), "\n")
	ts, _ = strconv.Atoi(array[0])
	for _, id := range strings.Split(array[1], ",") {
		nr, _ := strconv.Atoi(id)
		busses = append(busses, nr)
	}
	return
}

func pt1(ts int, busses []int) int {
	mn := ts
	busID := -1
	for _, bus := range busses {
		if bus == 0 {
			continue
		}
		// trick to compute next start, using ints
		departure := bus*(ts/bus) + bus
		delta := departure - ts
		// fmt.Println("bus will start at:", departure, "now it is:", ts, "We will wait:", delta)
		if delta < mn {
			busID = bus
			mn = delta
		}
	}
	fmt.Println("Bus:", busID, "will start in", mn, "minutes")
	return busID * mn
}

func pt2(ts int, busses []int) (current uint64) {
	incr := uint64(busses[0])
	// return uint64(LCM(17-4, 13-2, 19))
	parked := []park{}
	for i, bus := range busses[1:] {
		if bus != 0 {
			parked = append(parked, park{uint64(bus), uint64(i + 1)})
		}
	}
	ln := len(parked)
	current = 0
	fmt.Println(parked)
	for i := 0; i < ln; i++ {
		for (current+parked[i].offset)%parked[i].ID != 0 {
			current += incr
		}
		incr *= uint64(parked[i].ID)
	}
	return
}

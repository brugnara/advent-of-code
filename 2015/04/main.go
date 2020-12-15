package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/3

type point struct {
	x, y int
}

func main() {
	fmt.Println("PT1:", pt1("yzbqklnj"))
	fmt.Println("PT2:", pt2("yzbqklnj"))
}

func pt1(input string) (ret int) {
	start := 10
	for {
		if start > 10000000 {
			fmt.Println("break")
			break
		}
		h := md5.New()
		io.WriteString(h, input+strconv.Itoa(start))
		if strings.HasPrefix(fmt.Sprintf("%x", h.Sum(nil)), "00000") {
			return start
		}
		start++
	}
	return
}

func pt2(input string) (ret int) {
	start := 10
	for {
		if start > 10000000 {
			fmt.Println("break")
			break
		}
		h := md5.New()
		io.WriteString(h, input+strconv.Itoa(start))
		if strings.HasPrefix(fmt.Sprintf("%x", h.Sum(nil)), "000000") {
			return start
		}
		start++
	}
	return
}

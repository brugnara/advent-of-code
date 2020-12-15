package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2020/day/14

func main() {
	fmt.Println("PT1:", pt1(getInput()))
	fmt.Println("PT2:", pt2(getInput()))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

func pt1(lines []string) (ret int64) {
	reMask := regexp.MustCompile(`mask = (\w+)`)
	reMemory := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	hash := map[int]int64{}
	currentMask := ""
	for _, line := range lines {
		if line == "" {
			break
		}
		matchMask := reMask.FindStringSubmatch(line)
		if len(matchMask) != 0 {
			// fmt.Println("new mask:", matchMask[1])
			currentMask = matchMask[1]
			continue
		}
		matchMem := reMemory.FindStringSubmatch(line)
		address, _ := strconv.Atoi(matchMem[1])
		value, _ := strconv.Atoi(matchMem[2])
		bin := LeftPad2Len(strconv.FormatInt(int64(value), 2), "0", 36)
		merged := merge(bin, currentMask)
		hash[address], _ = strconv.ParseInt(merged, 2, 64)
		//
	}
	for _, v := range hash {
		ret += v
	}
	return
}

func merge(bin, mask string) (ret string) {
	tmp := []rune(bin)
	for i, s := range mask {
		if s == 'X' {
			continue
		}
		tmp[i] = s
	}
	return string(tmp)
}

func merge2(bin, mask string) []string {
	tmp := []rune(bin)
	xs := []int{}
	for i, s := range mask {
		switch s {
		case 'X':
			xs = append(xs, i)
			tmp[i] = 'X'
		case '1':
			tmp[i] = '1'
		}
	}
	return combo(string(tmp), xs, 0)
}

func combo(bin string, xs []int, index int) (ret []string) {
	if index == len(xs) {
		return []string{bin}
	}
	rn := []rune(bin)
	rn[xs[index]] = '0'
	ret = append(ret, combo(string(rn), xs, index+1)...)
	rn[xs[index]] = '1'
	ret = append(ret, combo(string(rn), xs, index+1)...)
	return
}

// LeftPad2Len https://github.com/DaddyOh/golang-samples/blob/master/pad.go
func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func pt2(lines []string) (ret int64) {
	reMask := regexp.MustCompile(`mask = (\w+)`)
	reMemory := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	hash := map[string]int64{}
	currentMask := ""
	for _, line := range lines {
		if line == "" {
			break
		}
		matchMask := reMask.FindStringSubmatch(line)
		if len(matchMask) != 0 {
			// fmt.Println("new mask:", matchMask[1])
			currentMask = matchMask[1]
			continue
		}

		matchMem := reMemory.FindStringSubmatch(line)
		address, _ := strconv.Atoi(matchMem[1])
		value, _ := strconv.Atoi(matchMem[2])
		bin := LeftPad2Len(strconv.FormatInt(int64(address), 2), "0", 36)
		merged := merge2(bin, currentMask)

		for _, addr := range merged {
			hash[addr] = int64(value)
		}
		//
	}
	for _, v := range hash {
		ret += v
	}
	return
}

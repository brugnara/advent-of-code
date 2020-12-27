package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// https://adventofcode.com/2020/day/19

func main() {
	lines := getInput()

	fmt.Println("PT1:", pt1(lines))
	fmt.Println("PT2:", pt2(lines))
}

func getInput() []string {
	raw, _ := ioutil.ReadFile("./input.txt")
	return strings.Split(string(raw), "\n")
}

func pt1(lines []string) (ret int) {
	// rules := NewRules()
	index := 0
	hash := map[string]string{}
	cache := map[string]string{}
	re := regexp.MustCompile(`(\d+): (.*)`)

	for i, line := range lines {
		if line == "" {
			// end of rules
			index = i + 1
			break
		}
		// rules.Add(line)
		match := re.FindStringSubmatch(line)
		switch match[2] {
		case "\"a\"", "\"b\"":
			cache[match[1]] = strings.ReplaceAll(match[2], "\"", "")
		default:
			hash[match[1]] = match[2]
		}
	}
	fmt.Println(index, cache)
	replacer := regexp.MustCompile(`(\d+)`)
	count := 0
	for {
		if len(hash) == 0 || count > 100 {
			fmt.Println(count)
			break
		}
		count++
		for k, v := range hash {
			tmp := replacer.ReplaceAllStringFunc(v, func(s string) string {
				if vv, ok := cache[s]; ok {
					if vv != "a" && vv != "b" {
						return "(" + vv + ")"
					}
					return vv
				}
				return s
			})
			if !replacer.MatchString(tmp) {
				cache[k] = tmp
				delete(hash, k)
			} else {
				hash[k] = tmp
			}
		}
	}
	fmt.Println(len(cache))
	reZero, _ := regexp.Compile("^" + strings.ReplaceAll(cache["0"], " ", "") + "$")
	fmt.Println(reZero)
	//
	for _, line := range lines[index:] {
		if reZero.MatchString(line) {
			ret++
		}
	}
	return
}

func pt2(lines []string) (ret int) {

	// rules := NewRules()
	index := 0
	hash := map[string]string{}
	cache := map[string]string{}
	re := regexp.MustCompile(`(\d+): (.*)`)

	for i, line := range lines {
		if line == "" {
			// end of rules
			index = i + 1
			break
		}
		match := re.FindStringSubmatch(line)
		switch match[2] {
		case "\"a\"", "\"b\"":
			cache[match[1]] = strings.ReplaceAll(match[2], "\"", "")
		default:
			hash[match[1]] = match[2]
		}
	}
	//
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	//
	// monkey patched here..
	hash["8"] = "42 +"
	hash["11"] = "42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31"
	//
	fmt.Println(index, cache)
	replacer := regexp.MustCompile(`(\d+)`)
	count := 0
	for {
		if len(hash) == 0 || count > 100 {
			fmt.Println(count)
			break
		}
		count++
		for k, v := range hash {
			tmp := replacer.ReplaceAllStringFunc(v, func(s string) string {
				if vv, ok := cache[s]; ok {
					if vv != "a" && vv != "b" {
						return "(" + vv + ")"
					}
					return vv
				}
				return s
			})
			if !replacer.MatchString(tmp) {
				cache[k] = tmp
				delete(hash, k)
			} else {
				hash[k] = tmp
			}
		}
	}
	fmt.Println(len(cache))
	reZero, _ := regexp.Compile("^" + strings.ReplaceAll(cache["0"], " ", "") + "$")
	fmt.Println(reZero)
	//
	for _, line := range lines[index:] {
		if reZero.MatchString(line) {
			ret++
		}
	}
	return
}

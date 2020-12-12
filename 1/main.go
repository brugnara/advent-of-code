package main

import "fmt"

func byTwo(input []int) {
	for i := 0; i < len(input)-1; i++ {
		for j := i; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				fmt.Println(input[i], input[j], input[i]+input[j], input[i]*input[j])
			}
		}
	}
}

func byThree(input []int) {
	for i := 0; i < len(input)-2; i++ {
		for j := i; j < len(input)-1; j++ {
			for z := j; z < len(input); z++ {
				if input[i]+input[j]+input[z] == 2020 {
					fmt.Println(input[i], input[j], input[z], input[i]+input[j]+input[z], input[i]*input[j]*input[z])
				}
			}
		}
	}
}

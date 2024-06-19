package main

import "fmt"
import "os"

func day1Part1 (inputStr []byte) int {
	floor := 0
	for _, char := range inputStr {
		if char == '(' {
			floor += 1
		} else if char == ')' {
			floor -= 1
		}
	}
	return floor
}

func day1Part2 (inputStr []byte) int {
	floor := 0
	for pos, char := range inputStr {
		if char == '(' {
			floor += 1
		} else if char == ')' {
			floor -= 1
		}

		if floor == -1 {
			return pos + 1
		}
	}
	return -1
}

func main () {
	bytes, _ := os.ReadFile("input/1-input.txt")
	fmt.Println("Part 1:", day1Part1(bytes))
	fmt.Println("Part 2:", day1Part2(bytes))
}


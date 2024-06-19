package main

import "fmt"
import "os"
import "bytes"

func day2Part1 (inputStr []byte) int {
	return 1
}

func main () {
	rawInput, _ := os.ReadFile("input/2-input.txt")
	inputByLine := bytes.Split(rawInput, []byte("\n"))
	fmt.Print(inputByLine)

	fmt.Println("Part 1:")
}
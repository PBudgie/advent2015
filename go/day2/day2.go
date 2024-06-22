package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day2Part2 (inputDimensions [][3]int) int {
	totalRibbonLength := 0

	for i := 0; i < len(inputDimensions); i++ {
		// Find ribbon needed to wrap the present
		dimensionsSlice := inputDimensions[i][0:3]
		slices.Sort(dimensionsSlice)
		wrappingRibbon := 2*dimensionsSlice[0] + 2*dimensionsSlice[1]

		// Find ribbon needed to tie the bow
		bowRibbon := inputDimensions[i][0] * inputDimensions[i][1] * inputDimensions[i][2]

		totalRibbonLength += wrappingRibbon + bowRibbon
	}

	return totalRibbonLength
}

func day2Part1 (inputDimensions [][3]int) int {
	totalWrappingPaper := 0

	for i := 0; i < len(inputDimensions); i++ {
		l, w, h := inputDimensions[i][0], inputDimensions[i][1], inputDimensions[i][2]
		// Necessary wrapping paper dimensions
		necessaryArea := 2*w*l + 2*l*h + 2*h*w

		// Find slack
		dimensionsSlice := inputDimensions[i][0:3]
		slices.Sort(dimensionsSlice)
		slack := dimensionsSlice[0] * dimensionsSlice[1]
		
		// Add to total
		totalWrappingPaper += necessaryArea + slack

		// Debugging
		// fmt.Printf("Dimensions: %v, Sorted: %v, Necessary: %v, Slack: %v\n", inputDimensions[i], dimensionsSlice, necessaryArea, slack)
	}
	
	return totalWrappingPaper
}

func main () {
	rawInput, _ := os.ReadFile("input/2-input.txt")
	inputByLine := strings.Split(string(rawInput), "\n")

	// Massage data into arrays of 3 integers for each dimension
	var inputDimensions [][3]int
	for i := 0; i < len(inputByLine); i++ {
		boxStrDimensions := strings.Split(inputByLine[i], "x")
		if (len(boxStrDimensions) == 3) {
			var boxNumDimensions [3]int
			for dim := 0; dim < 3; dim++ {
				num, _ := strconv.Atoi(boxStrDimensions[dim])
				boxNumDimensions[dim] = num
			}
			inputDimensions = append(inputDimensions, boxNumDimensions)
		} else {
			fmt.Print("More than 3 dimensions: ", boxStrDimensions)
		}
	}

	fmt.Println("Part 1:", day2Part1(inputDimensions))
	fmt.Println("Part 2:", day2Part2(inputDimensions))
}
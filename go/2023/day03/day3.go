package main

import (
	"fmt"
	"regexp"
	"strconv"

	utils "advent/utils"
)

func main() {
	input := utils.LoadFileAsStringSlice("input.txt")

	fmt.Printf("Part 1: %v\n", a(input))
	fmt.Printf("Part 2: %v\n", b(input))
}

func a(input []string) int {
	sum := 0
	for y, s := range input {
		re := regexp.MustCompile("[0-9]+")
		indices := re.FindAllStringIndex(s, -1)
		for _, index := range indices {

			x1, x2 := getXWindow(index[0], index[1]-index[0], len(s))

			if isPartNumber(x1, x2, y, input) {
				nInt, _ := strconv.Atoi(s[index[0]:index[1]])
				sum += nInt
			}
		}
	}

	return sum
}

func getXWindow(x int, wordLength int, rowLength int) (int, int) {
	var x1, x2 int

	x1 = x - 1
	if x1 < 0 {
		x1 = 0
	}

	x2 = x + wordLength + 1
	if x2 > rowLength {
		x2 = rowLength
	}

	return x1, x2
}

func getYWindow(y int, length int) (int, int) {
	var y1, y2 int

	y1 = y - 1
	if y1 < 0 {
		y1 = 0
	}

	y2 = y + 1
	if y2 > length {
		y2 = length - 1
	}
	return y1, y2
}

func isPartNumber(x1 int, x2 int, y int, input []string) bool {
	y1, y2 := getYWindow(y, len(input))

	result := false
	for i := y1; i <= y2 && i < len(input); i++ {

		stringToCheck := input[i][x1:x2] // slicing [inclusive : exclusive]
		found, _ := regexp.MatchString(`[^\d\.]`, stringToCheck)

		if found == true {
			result = true
		}
	}

	return result
}

type Coordinates struct {
	x, y int
}

func b(input []string) int {
	gearStore := map[Coordinates][]int{}

	for y, s := range input {
		re := regexp.MustCompile("[0-9]+")
		indices := re.FindAllStringIndex(s, -1)
		for _, index := range indices {

			x1, x2 := getXWindow(index[0], index[1]-index[0], len(s))

			x_g, y_g, found := potentialGearCoordinates(x1, x2, y, input)
			if found {
				nInt, _ := strconv.Atoi(s[index[0]:index[1]])
				coords := Coordinates{x_g, y_g}

				if _, ok := gearStore[coords]; ok {
					gearStore[coords] = append(gearStore[coords], nInt)
				} else {
					gearStore[coords] = []int{nInt}
				}
			}
		}
	}
	sum := 0

	for _, v := range gearStore {
		if len(v) > 1 {
			sum += v[0] * v[1]
		}
	}

	return sum
}

// this doesn't take into account if more than one * symbol was placed around a
// number but works for the test cases :)
func potentialGearCoordinates(x1 int, x2 int, y int, input []string) (int, int, bool) {
	y1, y2 := getYWindow(y, len(input))
	re := regexp.MustCompile(`\*`)

	for i := y1; i <= y2 && i < len(input); i++ {

		stringToCheck := input[i][x1:x2] // slicing [inclusive : exclusive]
		found := re.FindStringIndex(stringToCheck)

		if found != nil {
			return i, x1 + found[0], true
		}
	}

	return 0, 0, false
}

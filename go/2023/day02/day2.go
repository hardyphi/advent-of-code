package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "advent/utils"
)

func main() {
	input := utils.LoadFileAsStringSlice("input.txt")

	fmt.Printf("Part 1: %v\n", a(input))
	fmt.Printf("Part 2: %v\n", b(input))
}

func a(input []string) int {
	sum := 0
	for _, s := range input {
		split := strings.Split(s, ":")
		var gameId int
		_, _ = fmt.Sscanf(split[0], "Game %d", &gameId)
		// trying out Sscanf after the fact, definitely interesting as a way
		// of parsing very well understood data like this :)

		games := strings.Split(split[1], ";")

		canPlay := true
	gameLoop:
		for _, g := range games {
			drawn := strings.Split(g, ",")
			for _, balls := range drawn {
				var num int
				var colour string

				_, _ = fmt.Sscanf(balls, "%d %s", &num, &colour)

				if num > colourMap[colour] {
					canPlay = false
					break gameLoop
				}
			}
		}
		if canPlay == true {
			sum += gameId
		}
	}

	return sum
}

var colourMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func b(input []string) int {
	sum := 0
	for _, s := range input {
		split := strings.Split(s, ":")
		games := strings.Split(split[1], ";")

		colourStore := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, g := range games {
			drawn := strings.Split(g, ",")
			for _, balls := range drawn {
				split2 := strings.Split(strings.TrimSpace(balls), " ")
				colour := split2[1]
				num, _ := strconv.Atoi(split2[0])
				if num > colourStore[colour] {
					colourStore[colour] = num
				}
			}
		}
		var power int
		for _, v := range colourStore {
			if power == 0 {
				power = v
			} else {
				power *= v
			}

		}
		sum += power

	}

	return sum
}

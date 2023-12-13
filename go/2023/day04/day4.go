package main

import (
	"fmt"
	"slices"
	"strings"

	utils "advent/utils"
)

var splitFn = func(c rune) bool {
	return c == ' '
}

func main() {
	input := utils.LoadFileAsStringSlice("input.txt")

	fmt.Printf("Part 1: %v\n", a(input))
	fmt.Printf("Part 2: %v\n", b(input))
}

func a(input []string) int {
	total := 0
	for _, s := range input {
		score := 0
		ff := strings.Split(s, ":")
		ff2 := strings.Split(ff[1], "|")

		winners := strings.FieldsFunc(strings.TrimSpace(ff2[0]), splitFn)
		numbers := strings.FieldsFunc(strings.TrimSpace(ff2[1]), splitFn)

		for _, n := range numbers {
			if slices.Contains(winners, n) {
				if score == 0 {
					score++
				} else {
					score *= 2
				}
			}
		}
		total += score
	}
	return total
}

func b(input []string) int {
	scorecards := map[int]int{}

	for _, s := range input {
		ff := strings.Split(s, ":")
		ff2 := strings.Split(ff[1], "|")

		var card int
		_, _ = fmt.Sscanf(ff[0], "Card %d", &card)
		if _, ok := scorecards[card]; ok {
			scorecards[card]++
		} else {
			scorecards[card] = 1
		}
		winners := strings.FieldsFunc(strings.TrimSpace(ff2[0]), splitFn)
		numbers := strings.FieldsFunc(strings.TrimSpace(ff2[1]), splitFn)

		wins := 0
		for _, n := range numbers {
			if slices.Contains(winners, n) {
				wins++
			}
		}

		for i := card + 1; i <= card+wins && i <= len(input); i++ {
			if _, ok := scorecards[i]; ok {
				scorecards[i] += scorecards[card]
			} else {
				scorecards[i] = scorecards[card]
			}
		}

	}
	total := 0
	for _, v := range scorecards {
		total += v
	}

	return total
}

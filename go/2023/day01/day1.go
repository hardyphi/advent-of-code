package main

import (
	"fmt"
	"regexp"
	"strconv"

	utils "advent/utils"
)

var wordToNumber = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	input := utils.LoadFileAsStringSlice("input.txt")

	fmt.Printf("Part 1: %v\n", a(input))
	fmt.Printf("Part 2: %v\n", b(input))
}

func a(input []string) int {
	sum := 0
	for _, s := range input {
		re := regexp.MustCompile("[0-9]")
		filtered := re.FindAllString(s, -1)
		combined := filtered[0] + filtered[len(filtered)-1]
		number, err := strconv.Atoi(combined)
		if err != nil {
			panic(err)
		}
		sum += number
	}
	return sum
}

func b(input []string) int {
	strings := ""
	for k := range wordToNumber {
		strings += "|" + k
	}

	expr := "[0-9]" + strings

	sum := 0
	for _, s := range input {
		re := regexp.MustCompile(expr)
		first := re.FindString(s)
		last := ""
		for i := len(s) - 1; i >= 0; i-- {
			last = re.FindString(s[i:])
			if last != "" {
				break
			}
		}

		number, err := strconv.Atoi(getFirstAndLastConverted(first, last))
		if err != nil {
			panic(err)
		}
		sum += number
	}

	return sum
}

func getFirstAndLastConverted(first string, last string) string {
	return mapWordToNumber(first) + mapWordToNumber(last)

}

func mapWordToNumber(s string) string {
	if wordToNumber[s] != "" {
		return wordToNumber[s]
	}
	return s

}

package main

import (
	"fmt"
	"strconv"

	utils "advent/utils"
)

func main() {
	input := utils.LoadFileAsStringSlice("input.txt")

	fmt.Printf("Power Consumption = %v\n", powerConsumption(input))
	fmt.Printf("Life Support = %v\n", lifeSupport(input))
}

func powerConsumption(input []string) int64 {
	countOnes := counter(input)
	var gammaString string
	var epsiString string
	for _, count := range countOnes {
		if mostCommon(count ,len(input)) {
			gammaString += "1"
			epsiString += "0"
		} else {
			gammaString += "0"
			epsiString += "1"
		}
	}

	gamma := bToi(gammaString)
	epsi := bToi(epsiString)
	return gamma * epsi
}

func lifeSupport(input []string) int64 {
	oxygenRating := input
	co2Rating := input
	oxygenString := filterReport(oxygenRating, mostCommon)
	co2String := filterReport(co2Rating, leastCommon)

	oxygen := bToi(oxygenString)
	co2 := bToi(co2String)
	return oxygen  * co2
}

func counter(slice []string) []int {
	length := len(slice[0])
	count := make([]int, length) //had a lot of fun learning about slices here...
	for _, s := range slice {
		for i, c := range s {
			if string(c) == "1" {
				count[i]++
			}
		}
	}
	return count
}

func filterReport(reportInput []string, conditional func(int, int) bool) string {
	report := make([]string, len(reportInput))
	copy(report, reportInput)

	for i := 0; len(report) > 1; i++ {
		countOnes := counter(report)
		if conditional(countOnes[i], len(report)) {
			report = removeMatching(report, "0", i)
		} else {
			report = removeMatching(report, "1", i)
		}
	}
	return report[0]
}

func mostCommon(count int, length int) bool {
	return count >= length - count
}

func leastCommon(count int, length int) bool {
	return count < length - count
}

func removeMatching(slice []string, value string, index int) []string {
	for j := 0; j < len(slice); j++ {
		if string(slice[j][index]) == value {
			if len(slice) == 1 {
				break
			}
			slice = append(slice[:j], slice[j+1:]...)
			j--
		}
	}
	return slice
}

func bToi(b string) int64 {
	i, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic(err)
	}
	return i
}

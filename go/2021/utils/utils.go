package utils

import (
	"bufio"
	"os"
	"strconv"
)

func loadFile(filepath string) []string {

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			panic(err)
		}
		result = append(result, line)
	}

	return result
}

func LoadFileAsStringSlice(filepath string) []string {
    return loadFile(filepath)
}

func LoadFileAsIntSlice(filepath string) []int {
    strings := loadFile(filepath)
    var ints = []int{}
    for _, s := range strings {
        int, err := strconv.Atoi(s)
        if err != nil {
            panic(err)
        }
        ints = append(ints, int)
    }
    return ints
}


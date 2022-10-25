package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "advent/utils"
)

func main() {
    input := utils.LoadFileAsStringSlice("input.txt")
    
    fmt.Printf("a = %v\nb = %v\n", a(input), b(input))
}

func a(input []string) int {
    x := 0
    depth := 0
    
    for _, s := range input {
        vector := strings.Split(s, " ")
        magnitude, err := strconv.Atoi(vector[1])
        if err != nil {
            panic(err)
        }
        switch vector[0] {
        case "forward":
            x += magnitude    
        case "up":
            depth -= magnitude
        case "down":
            depth += magnitude
        }
    }
    return x * depth
}

func b(input []string) int {
    x := 0
    depth := 0
    aim := 0
    
    for _, s := range input {
        vector := strings.Split(s, " ")
        magnitude, err := strconv.Atoi(vector[1])
        if err != nil {
            panic(err)
        }
        switch vector[0] {
        case "forward":
            x += magnitude    
            depth += magnitude * aim
        case "up":
            aim -= magnitude
        case "down":
            aim += magnitude
        }
    }
    return x * depth
}



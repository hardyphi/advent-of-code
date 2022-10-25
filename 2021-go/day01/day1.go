package main

import (
    "fmt"

    utils "advent/utils"
)

func main() {
    input := utils.LoadFileAsIntSlice("input.txt")

    resultA := a(input)
    resultB := b(input)
    fmt.Printf("A: %v\nB: %v\n", resultA, resultB)
}

func a(input []int) int {
    result := 0
    for i, s := range input {
        if i==0 {
        } else if s - input[i-1] > 0 {
            result++
        }
    }
    return result
}

func b(input []int) int {
    result := 0
    for i, s := range input {
        if i < 3 {
        }  else if s - input[i-3] > 0 {
            result++
        }
    }   
    return result
}

package main

import (
	"fmt"
	"strings"
    "strconv"

	utils "advent/utils"
)

func main() {
	input := utils.LoadFileAsStringSlice("input.txt")

	var numbers []string
	var b [][]string
	var boards [][][]string
	for i, l := range input {
		if i == 0 {
			numbers = strings.Split(l, ",")
			continue
		}
		if i == 1 {
			continue
		}
		if len(l) == 0 {
			boards = append(boards, b)
			b = [][]string{}
            continue
		}
		boxes := strings.Fields(l)
		b = append(b, boxes)
	}
	set := NewSet()
    
    var winner int
    var final int
    game:
	for _, num := range numbers {
		set.Add(num)
		for i, board := range boards {
			if (checkHorizontal(board, set) || checkVertical(board, set)) {
                final, _ = strconv.Atoi(num) 
                winner = i
                break game
            }
		}
	}
    
    fmt.Printf("The winner is %v\n", winner)
    winningBoard := boards[winner]
    sum := 0
    for _, row := range winningBoard {
        for _, s := range row {
            if !set.Has(s) {
                num,_ := strconv.Atoi(s)
                sum += num
                fmt.Println(sum)
            }
        }
    }
    score := sum * final 
    fmt.Printf("With a score of %v\n", score)
}

type Set struct {
	list map[string]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.list = make(map[string]struct{})
	return s
}

func (s *Set) Add(v string) {
	s.list[v] = struct{}{}
}

func (s *Set) Has(v string) bool {
	_, ok := s.list[v]
	return ok
}

func checkHorizontal(board [][]string, set *Set) bool {
	for _, row := range board {
		count := 0
		for _, s := range row {
			if set.Has(s) {
				count++
			}
			if count == len(board) {
				return true
			}
		}
	}
	return false
}

func checkVertical(board [][]string, set *Set) bool {
	for i := 0; i < len(board); i++ {
		count := 0
		for j := 0; j < len(board); j++ {
			if set.Has(board[j][i]) {
				count++
			}
			if count == len(board) {
				return true
			}
		}
	}
	return false
}

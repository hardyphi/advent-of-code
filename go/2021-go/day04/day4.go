package main

import (
	"fmt"
	"strconv"
	"strings"

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

	winner, winNum, winSet := winGame(boards, numbers)
	winningScore := scoreOfBoard(boards[winner], winSet) * winNum
	fmt.Printf("Winning Board is %v\nWith a score of %v\n", winner, winningScore)

	loser, loseNum, loseSet := loseGame(boards, numbers)
	losingScore := scoreOfBoard(loser, loseSet) * loseNum
	fmt.Printf("The losing score is %v\n", losingScore)
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

func winGame(boards [][][]string, numbers []string) (winner int, winNum int, winSet *Set) {
	set := NewSet()
	for _, num := range numbers {
		set.Add(num)
		for i, board := range boards {
			if checkBoard(board, set) {
				last, _ := strconv.Atoi(num)
				return i, last, set
			}
		}
	}
	return 0, 0, nil
}

func loseGame(boards [][][]string, numbers []string) (loser [][]string, loseNum int, loseSet *Set) {
	b := make([][][]string, len(boards))
	copy(b, boards)
	set := NewSet()
	for _, num := range numbers {
		set.Add(num)
		for i := 0; i < len(b); i++ {
			if checkBoard(b[i], set) {
                if len(b) == 1 {
                    last,_ := strconv.Atoi(num)
                    return b[0], last, set
                }
				b = removeBoard(b, i)
				if i > 1 {
					i--
				}
			}
		}
	}
	return nil, 0, nil
}

func checkBoard(board [][]string, set *Set) bool {
	for i := 0; i < len(board); i++ {
		hCount := 0
		vCount := 0
		for j := 0; j < len(board); j++ {
			if set.Has(board[i][j]) {
				hCount++
			}
			if set.Has(board[j][i]) {
				vCount++
			}
			if vCount == len(board) || hCount == len(board) {
				return true
			}
		}
	}
	return false
}

func scoreOfBoard(board [][]string, set *Set) int {
	sum := 0
	for _, row := range board {
		for _, s := range row {
			if !set.Has(s) {
				num, _ := strconv.Atoi(s)
				sum += num
			}
		}
	}
	return sum
}

func removeBoard(b [][][]string, i int) [][][]string {
	return append(b[:i], b[i+1:]...)
}

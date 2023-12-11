package main

import "testing"

func Test_A1(t *testing.T) {
	test := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}
	got := a(test)

	want := 4361

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func Test_A2(t *testing.T) {
	test := []string{
		"12.......*..",
		"+.........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78..........",
		".......23...",
		"....90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1.......56"}
	got := a(test)

	want := 413

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func Test_A3(t *testing.T) {
	test := []string{
		"12.......*..",
		"+.........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78.........9",
		".5.....23..$",
		"8...90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1..503+.56"}
	got := a(test)

	want := 925

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func Test_A4(t *testing.T) {
	test := []string{
		".......5......",
		"..7*..*.......",
		"...*13*.......",
		".......15....."}
	got := a(test)

	want := 40

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

// this was the big one... looking for 4 in the string
func Test_A5(t *testing.T) {
	test := []string{"........", ".24..4..", "......*."}
	got := a(test)

	want := 4

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestB1(t *testing.T) {
	test := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}
	got := b(test)

	want := 467835

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

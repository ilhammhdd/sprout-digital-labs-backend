package main

import "fmt"

var board = [10][10]string{
	{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
	{"8", "RB", "TB", "BB", "QB", "KB", "BB", "TB", "RB", "8"},
	{"7", "PB", "PB", "PB", "PB", "PB", "PB", "PB", "PB", "7"},
	{"6", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "6"},
	{"5", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "5"},
	{"4", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "4"},
	{"3", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "3"},
	{"2", "PW", "PW", "PW", "PW", "PW", "PW", "PW", "PW", "2"},
	{"1", "RW", "TW", "BW", "QW", "KW", "BW", "TW", "RW", "1"},
	{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
}

type turn struct {
	id          uint16
	playerWhite bool
	action      [2]string
}

var turns []turn

var captured [2][]string

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (i+j)%2 != 0 && board[i][j] == "  " {
				fmt.Print(" ## ")
			} else if j == 0 || j == 9 {
				fmt.Printf("%s", board[i][j])
			} else {
				fmt.Printf(" %s ", board[i][j])
			}
		}
		fmt.Println()
	}
}

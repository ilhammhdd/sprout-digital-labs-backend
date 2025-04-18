package state

import "github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"

var Board = [10][10]string{
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

func getBoardIndices(square entity.Square) entity.Indices {
	if square[0] >= 'a' && square[0] <= 'z' {
		square[0] -= 'a' - 'A'
	}
	return entity.Indices{int(9 - (square[1] - '0')), int(square[0] - 'A' + 1)}
}

package state

import "github.com/ilhammhdd/sprout-digital-labs-backend/internal/entity"

var Board = [10][10]string{
	{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
	{"8", "RD", "TD", "BD", "QD", "KD", "BD", "TD", "RD", "8"},
	{"7", "PD", "PD", "PD", "PD", "PD", "PD", "PD", "PD", "7"},
	{"6", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "6"},
	{"5", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "5"},
	{"4", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "4"},
	{"3", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "  ", "3"},
	{"2", "PL", "PL", "PL", "PL", "PL", "PL", "PL", "PL", "2"},
	{"1", "RL", "TL", "BL", "QL", "KL", "BL", "TL", "RL", "1"},
	{" ", "A ", "B ", "C ", "D ", "E ", "F ", "G ", "H ", " "},
}

func GetBoardIndices(square entity.Square) entity.Indices {
	if square[0] >= 'a' && square[0] <= 'z' {
		square[0] -= 'a' - 'A'
	}
	return entity.Indices{int(9 - (square[1] - '0')), int(square[0] - 'A' + 1)}
}

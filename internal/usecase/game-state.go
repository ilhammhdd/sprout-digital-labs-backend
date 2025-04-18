package usecase

import (
	"fmt"

	"github.com/ilhammhdd/sprout-digital-labs-backend/internal/state"
)

func PrintState() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (i+j)%2 != 0 && state.Board[i][j] == "  " {
				fmt.Print(" ## ")
			} else if j == 0 || j == 9 {
				fmt.Printf("%s", state.Board[i][j])
			} else {
				fmt.Printf(" %s ", state.Board[i][j])
			}
		}
		fmt.Println()
	}
}

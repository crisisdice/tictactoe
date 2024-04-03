package rendering

import (
	"fmt"

	. "example/tictactoe/internal/types"
)

const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Purple = "\033[35m"
)

func GetPlayerSymbol(player Player) string {
	switch player {
	case X:
		{
			return "X"
		}
	case O:
		{
			return "O"
		}
	}

	panic("player was not initialized")
}

func InitBoard() Board {
	var board = Board{
		{
			"",
			"",
			"",
		},
		{
			"",
			"",
			"",
		},
		{
			"",
			"",
			"",
		},
	}

	return board
}

func PrintBoard(board Board, cursor Cursor, player Player, printPlayer bool) {
	var (
		cursorX = cursor[0]
		cursorY = cursor[1]
	)

	fmt.Print("\n\n\n\n\n")

	fmt.Println(" _____   _            _____                    _____                ")
	fmt.Println("|_   _| (_)   ___    |_   _|   __ _    ___    |_   _|   ___     ___ ")
	fmt.Println("  | |   | |  / __|     | |    / _` |  / __|     | |    / _ \\   / _ \\")
	fmt.Println("  | |   | | | (__      | |   | (_| | | (__      | |   | (_) | |  __/")
	fmt.Println("  |_|   |_|  \\___|     |_|    \\__,_|  \\___|     |_|    \\___/   \\___|")

	fmt.Print("\n\n")

	fmt.Print("Press Esc or q to quit; use arrows or h,j,k,l to move; Space to enter move.")

	fmt.Print("\n\n")

	for rowIndex, row := range board {
		fmt.Print("\t")
		for collumnIndex, collumn := range row {
			if collumnIndex == cursorX && rowIndex == cursorY && printPlayer {
				fmt.Print(Purple, GetPlayerSymbol(player), Reset)
			} else {
				fmt.Print(Green, collumn, Reset)

				if collumn == "" {
					fmt.Print(" ")
				}
			}

			if collumnIndex < 2 {
				fmt.Print("|")
			}

		}

		if rowIndex < 2 {
			fmt.Print("\n\t-+-+-")
		}

		fmt.Print("\n")
	}

	fmt.Print("\n")
}

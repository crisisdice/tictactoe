package main

import (
	"fmt"
	"os"

	. "example/tictactoe/internal/types"

	"example/tictactoe/internal/input"
	"example/tictactoe/internal/opponent"
	"example/tictactoe/internal/rendering"
	"example/tictactoe/internal/rules"

	"github.com/eiannone/keyboard"
)

func main() {
	var (
		board               = rendering.InitBoard()
		cursor              = Cursor{1, 0}
		activePlayer Player = X
		winner       Player
		turnCount    = 0
	)

	for {
		// clear screen
		fmt.Print("\033[H\033[2J")

		rules.DetectWin(board, &winner)

		var (
			gameHasWinner      = winner != 0
			gameIsDraw         = rules.DetectDraw(board)
			renderPlayerCursor = !gameHasWinner && !gameIsDraw && activePlayer == O
		)

		rendering.PrintBoard(board, cursor, activePlayer, renderPlayerCursor)

		if gameHasWinner {
			fmt.Println("Winner: ", rendering.GetPlayerSymbol(winner), "!!!")
			os.Exit(0)
		}

		if gameIsDraw {
			fmt.Println("It's a draw.")
			os.Exit(0)
		}

		if activePlayer == X {
			opponent.ComputerMove(&board, turnCount, activePlayer)

			turnCount += 1

			rules.SwitchActivePlayer(&activePlayer)

			continue
		}

		char, key, err := keyboard.GetSingleKey()

		if err != nil {
			panic(err)
		}

		input.HandleMove(char, key, &board, &cursor, &activePlayer)
	}
}

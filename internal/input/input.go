package input

import (
	"os"

	. "example/tictactoe/internal/types"

	"example/tictactoe/internal/rendering"
	"example/tictactoe/internal/rules"

	"github.com/eiannone/keyboard"
)

func HandleMove(char rune, key keyboard.Key, board *Board, cursor *Cursor, activePlayer *Player) {
	var (
		cursorX = cursor[0]
		cursorY = cursor[1]
	)

	switch {
	case key == keyboard.KeyEsc || char == 'q':
		{
			os.Exit(0)
		}
	case key == keyboard.KeySpace:
		{
			var isLegalMove = rules.IsLegalMove(board, cursor)

			if isLegalMove {
				board[cursorY][cursorX] = rendering.GetPlayerSymbol(*activePlayer)
				rules.SwitchActivePlayer(activePlayer)
			}
		}
	case char == 'h' || key == keyboard.KeyArrowLeft:
		{
			var cursorX = (cursorX + 5) % 3
			*cursor = Cursor{cursorX, cursorY}
		}
	case char == 'j' || key == keyboard.KeyArrowDown:
		{
			var cursorY = (cursorY + 1) % 3
			*cursor = Cursor{cursorX, cursorY}
		}
	case char == 'k' || key == keyboard.KeyArrowUp:
		{
			var cursorY = (cursorY + 5) % 3
			*cursor = Cursor{cursorX, cursorY}
		}
	case char == 'l' || key == keyboard.KeyArrowRight:
		{
			var cursorX = (cursorX + 1) % 3
			*cursor = Cursor{cursorX, cursorY}
		}
	}
}

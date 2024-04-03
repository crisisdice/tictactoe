package opponent

import (
	. "example/tictactoe/internal/types"

	"example/tictactoe/internal/rendering"
	"example/tictactoe/internal/rules"
)

func detectOddSymbolOut(unit [3]string, one string, two string) int {
	if unit[1] == one && unit[1] == unit[2] && unit[0] == two {
		return 0
	}

	if unit[0] == one && unit[0] == unit[2] && unit[1] == two {
		return 1
	}

	if unit[0] == one && unit[0] == unit[1] && unit[2] == two {
		return 2
	}

	return -1
}

func detectAlmostWin(unit [3]string, player Player) int {
	var playerSymbol = rendering.GetPlayerSymbol(player)
	return detectOddSymbolOut(unit, playerSymbol, "")
}

func playToAlmostWin(board *Board, player Player, playToMine bool) bool {
	var (
		searchingFor                                Player
		playerSymbol                                = rendering.GetPlayerSymbol(player)
		rows, collumns, leftDiagonal, rightDiagonal = rules.SplitBoard(board)
	)

	if playToMine {
		searchingFor = player
	} else {
		searchingFor = rules.GetOtherPlayer(player)
	}

	for rowIndex, row := range rows {
		var possibleWin = detectAlmostWin(row, searchingFor)

		if possibleWin > -1 {
			board[rowIndex][possibleWin] = playerSymbol

			return true
		}
	}

	for collumnIndex, collumn := range collumns {
		var possibleWin = detectAlmostWin(collumn, searchingFor)

		if possibleWin > -1 {
			board[possibleWin][collumnIndex] = playerSymbol

			return true
		}
	}

	var leftDiagonalWin = detectAlmostWin(leftDiagonal, searchingFor)

	if leftDiagonalWin > -1 {
		board[leftDiagonalWin][leftDiagonalWin] = playerSymbol

		return true
	}

	var rightDiagonalWin = detectAlmostWin(rightDiagonal, searchingFor)

	if rightDiagonalWin > -1 {
		switch rightDiagonalWin {
		case 0:
			{
				board[0][2] = playerSymbol
			}
		case 1:
			{
				board[1][1] = playerSymbol
			}
		case 2:
			{
				board[2][0] = playerSymbol
			}
		}

		return true
	}

	return false

}

func blockOtherPlayer(board *Board, player Player) bool {
	return playToAlmostWin(board, player, false)
}

func takeUnblockedWin(board *Board, player Player) bool {
	return playToAlmostWin(board, player, true)
}

func detectOpening(board *Board, player Player) bool {
	var rows, collumns, _, _ = rules.SplitBoard(board)

	var playerSymbol = rendering.GetPlayerSymbol(player)

	for rowIndex, row := range rows {
		var lastChance = detectOddSymbolOut(row, "", playerSymbol)

		if lastChance == 2 {
			board[rowIndex][0] = rendering.GetPlayerSymbol(player)

			return true
		}
	}

	for collumnIndex, collumn := range collumns {
		var lastChance = detectOddSymbolOut(collumn, "", playerSymbol)

		if lastChance == 2 {
			board[0][collumnIndex] = playerSymbol

			return true
		}
	}
	return false
}

func playToDraw(board *Board, player Player) {
	if blockOtherPlayer(board, player) {
		return
	}

	if detectOpening(board, player) {
		return
	}

	for rowIndex, row := range board {
		for collumnIndex, collumn := range row {
			if collumn == "" {
				board[rowIndex][collumnIndex] = rendering.GetPlayerSymbol(player)
				return
			}
		}
	}
}

func ComputerMove(board *Board, turnCount int, player Player) {
	if takeUnblockedWin(board, player) {
		return
	}

	var playerSymbol = rendering.GetPlayerSymbol(player)

	switch turnCount {
	case 0:
		{
			board[0][0] = playerSymbol
			return
		}
	case 1:
		{
			if board[1][1] != "" {
				board[2][2] = playerSymbol
				return
			}

			var (
				lowerPlayCases = board[1][0] != "" || board[2][0] != "" || board[2][1] != "" || board[2][2] != ""
				upperPlayCases = board[0][1] != "" || board[0][2] != "" || board[1][2] != ""
			)

			if lowerPlayCases {
				board[0][2] = playerSymbol
			}

			if upperPlayCases {
				board[2][0] = playerSymbol
			}

			return
		}
	case 2:
		{
			if board[1][1] != "" {
				if board[0][2] != "" {
					board[2][0] = playerSymbol
					return
				}

				if board[2][0] != "" {
					board[0][2] = playerSymbol
					return
				}

				playToDraw(board, player)

				return
			}

			if board[2][2] != "" {
				board[2][0] = playerSymbol
				return
			}

			board[2][2] = playerSymbol
		}
	default:
		{
			playToDraw(board, player)
		}
	}
}
